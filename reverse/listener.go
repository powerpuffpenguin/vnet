package reverse

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/powerpuffpenguin/vnet"
)

type dialResult struct {
	c net.Conn
	e error
}
type Listener struct {
	opts listenerOptions
	addr net.Addr

	ch    chan dialResult
	close <-chan struct{}
	done  uint32
	m     sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc
}

func Listen(addr net.Addr, opt ...ListenerOption) *Listener {
	opts := defaultListenerOptions
	for _, o := range opt {
		o.apply(&opts)
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Listener{
		opts: opts,
		addr: addr,

		ch:     make(chan dialResult),
		close:  ctx.Done(),
		ctx:    ctx,
		cancel: cancel,
	}
}

// Accept waits for and returns the next connection to the listener.
func (l *Listener) Accept() (c net.Conn, e error) {
	// check closed
	if atomic.LoadUint32(&l.done) != 0 {
		e = vnet.ErrListenerClosed
		return
	}

	// dial
	go l.asyncDial()

	// wait result
	select {
	case result := <-l.ch:
		c = result.c
		e = result.e
	case <-l.close:
		e = vnet.ErrListenerClosed
	}
	return
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (l *Listener) Close() (e error) {
	if atomic.LoadUint32(&l.done) == 0 {
		l.m.Lock()
		defer l.m.Unlock()
		if l.done == 0 {
			defer atomic.StoreUint32(&l.done, 1)
			l.cancel()
			return
		}
	}
	e = vnet.ErrListenerClosed
	return
}

// Addr returns the listener's network address.
func (l *Listener) Addr() net.Addr {
	return l.addr
}
func (l *Listener) asyncDial() {
	c, e := l.dial()
	select {
	case <-l.close:
		if e == nil {
			c.Close()
		}
	case l.ch <- dialResult{
		c: c,
		e: e,
	}:
	}
}
func (l *Listener) dial() (c net.Conn, e error) {
	opts := &l.opts
	if opts.dialContext != nil {
		c, e = opts.dialContext(l.ctx, l.addr.Network(), l.Addr().String())
		if e != nil {
			return
		}
	} else if opts.dial != nil {
		c, e = opts.dial(l.addr.Network(), l.Addr().String())
		if e != nil {
			return
		}
		select {
		case <-l.close:
			c.Close()
			e = vnet.ErrListenerClosed
		default:
		}
	} else {
		// default dial tcp
		var d net.Dialer
		c, e = d.DialContext(l.ctx, l.addr.Network(), l.Addr().String())
		if e != nil {
			return
		}
	}

	if opts.synAck {
		ch := make(chan error, 1)
		go l.asyncSynAck(ch, c)
		select {
		case e = <-ch:
			if e != nil {
				c.Close()
				c = nil
			}
		case <-l.close:
			e = vnet.ErrListenerClosed
			c.Close()
			c = nil
		}
	}
	return
}

func (l *Listener) asyncSynAck(ch chan<- error, c net.Conn) {
	opts := &l.opts
	stream := &datagramStream{
		rw: c,
	}
	buffer := make(chan error, 1)
	// recv syn
	e := l.recvSyn(buffer, stream, opts.heartTimeout)
	if e != nil {
		ch <- e
		return
	}
	// send syn+ack
	// recv ack
	ch <- l.sendSynAck(buffer, stream, opts.synAckTimeout)
}
func (l *Listener) sendSynAck(ch chan error, stream *datagramStream, timeout time.Duration) (e error) {
	var t *time.Timer
	var deadline <-chan time.Time
	if timeout > 0 {
		t = time.NewTimer(timeout)
		deadline = t.C
	}
	go func() {
		var err error
		err = stream.Send(DatagramSynAck)
		if err == nil {
			err = stream.Recv(DatagramAck)
			if err == nil {
				close(ch)
				return
			}
		}
		ch <- err
	}()
	select {
	case <-deadline:
		e = context.DeadlineExceeded
		return
	case e = <-ch:
	}
	if t != nil && !t.Stop() {
		<-t.C
	}
	return
}
func (l *Listener) recvSyn(ch chan error, stream *datagramStream, timeout time.Duration) (e error) {
	var t *time.Timer
	var deadline <-chan time.Time
	var heart chan bool
	if timeout > 0 {
		t = time.NewTimer(timeout)
		deadline = t.C
		heart = make(chan bool, 1)
	}
	go func() {
		var err error
		for {
			err = stream.Recv(DatagramHeart, DatagramSyn)
			if err != nil {
				break
			} else if stream.Event() == DatagramSyn {
				break
			} else if heart != nil {
				select {
				case heart <- true:
				default:
				}
			}
		}
		ch <- err
	}()
	work := true
	for work {
		select {
		case <-deadline:
			e = context.DeadlineExceeded
			return
		case e = <-ch:
			work = false
		case <-heart:
			if !t.Stop() {
				<-t.C
			}
			t.Reset(timeout)
		}
	}
	if t != nil && !t.Stop() {
		<-t.C
	}
	return
}
