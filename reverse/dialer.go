package reverse

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/powerpuffpenguin/vnet"
)

type Dialer struct {
	opts dialerOptions
	l    net.Listener

	ch    chan *datagramStream
	close <-chan struct{}
	done  uint32
	m     sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc
}

func NewDialer(l net.Listener, opt ...DialerOption) *Dialer {
	opts := defaultDialerOptions
	for _, o := range opt {
		o.apply(&opts)
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Dialer{
		opts: opts,
		l:    l,

		ch:     make(chan *datagramStream),
		close:  ctx.Done(),
		ctx:    ctx,
		cancel: cancel,
	}
}
func (d *Dialer) Close() (e error) {
	if atomic.LoadUint32(&d.done) == 0 {
		d.m.Lock()
		defer d.m.Unlock()
		if d.done == 0 {
			defer atomic.StoreUint32(&d.done, 1)
			d.cancel()
			d.l.Close()
			return
		}
	}
	e = vnet.ErrDialerClosed
	return
}
func (d *Dialer) Serve() error {
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		c, e := d.l.Accept()
		if e != nil {
			select {
			case <-d.close:
				return vnet.ErrDialerClosed
			default:
			}
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		go d.onAccept(c)
	}
}
func (d *Dialer) onAccept(c net.Conn) {
	stream := &datagramStream{
		rw: c,
	}
	var deadline <-chan time.Time
	var t *time.Timer
	if d.opts.heart > 0 {
		t = time.NewTimer(d.opts.heart)
		deadline = t.C
	}
	work := true
	for work {
		select {
		case <-d.close:
			work = false
			c.Close()
		case <-deadline:
			e := d.sendHeart(stream)
			if e != nil {
				c.Close()
				return
			} else if t != nil {
				t.Reset(d.opts.heart)
			}
		case d.ch <- stream:
			work = false
		}
	}
	if t != nil && !t.Stop() {
		<-t.C
	}
}
func (d *Dialer) sendHeart(stream *datagramStream) (e error) {
	if d.opts.heartTimeout < 1 {
		e = stream.Send(DatagramHeart)
	} else {
		t := time.NewTimer(d.opts.heartTimeout)
		ch := make(chan error, 1)
		go func() {
			ch <- stream.Send(DatagramHeart)
		}()
		select {
		case <-d.close:
			e = vnet.ErrDialerClosed
		case e = <-ch:
		case <-t.C:
			e = context.DeadlineExceeded
			return
		}
		if !t.Stop() {
			<-t.C
		}
	}
	return
}
func (d *Dialer) Dial(network, addr string) (c net.Conn, e error) {
	return d.DialContext(context.Background(), network, addr)
}
func (d *Dialer) DialContext(ctx context.Context, network, addr string) (c net.Conn, e error) {
	var stream *datagramStream
	select {
	case <-ctx.Done():
		e = ctx.Err()
		return
	case stream = <-d.ch:
	case <-d.close:
		e = vnet.ErrDialerClosed
		return
	}

	if d.opts.synAck {
		e = d.synAck(ctx, stream)
		if e != nil {
			stream.rw.Close()
			return
		}
	}
	c = stream.rw
	return
}
func (d *Dialer) synAck(ctx context.Context, stream *datagramStream) (e error) {
	var t *time.Timer
	var deadline <-chan time.Time
	if d.opts.timeout > 0 {
		t = time.NewTimer(d.opts.timeout)
		deadline = t.C
	}
	ch := make(chan error, 1)
	go d.asyncSynAck(ch, stream)
	select {
	case <-ctx.Done():
		e = ctx.Err()
	case <-deadline:
		e = context.DeadlineExceeded
		return
	case e = <-ch:
	case <-d.close:
		e = vnet.ErrDialerClosed
	}
	if t != nil && !t.Stop() {
		<-t.C
	}
	return
}
func (d *Dialer) asyncSynAck(ch chan<- error, stream *datagramStream) {
	// dial send syn
	e := stream.Send(DatagramSyn)
	if e != nil {
		ch <- e
		return
	}
	// recv syn+ack
	e = stream.Recv(DatagramSynAck)
	if e != nil {
		ch <- e
		return
	}
	// send ack
	e = stream.Send(DatagramAck)
	if e != nil {
		ch <- e
		return
	}
}
