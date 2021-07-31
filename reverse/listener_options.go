package reverse

import (
	"context"
	"net"
	"time"
)

var defaultListenerOptions = listenerOptions{
	dial:          nil,
	dialContext:   nil,
	synAck:        true,
	synAckTimeout: time.Second * 75,
	heartTimeout:  time.Second * 75,
}

type listenerOptions struct {
	dial          func(network, address string) (net.Conn, error)
	dialContext   func(ctx context.Context, network, addr string) (conn net.Conn, e error)
	synAck        bool
	synAckTimeout time.Duration
	heartTimeout  time.Duration
}

type ListenerOption interface {
	apply(*listenerOptions)
}
type funcListenerOption struct {
	f func(*listenerOptions)
}

func (fdo *funcListenerOption) apply(do *listenerOptions) {
	fdo.f(do)
}
func newListenerOption(f func(*listenerOptions)) *funcListenerOption {
	return &funcListenerOption{
		f: f,
	}
}
func WithListenerDial(f func(network, address string) (net.Conn, error)) ListenerOption {
	return newListenerOption(func(o *listenerOptions) {
		o.dial = f
	})
}
func WithListenerDialContext(f func(ctx context.Context, network, address string) (net.Conn, error)) ListenerOption {
	return newListenerOption(func(o *listenerOptions) {
		o.dialContext = f
	})
}
func WithListenerSynAck(synAck bool) ListenerOption {
	return newListenerOption(func(o *listenerOptions) {
		o.synAck = synAck
	})
}
func WithListenerSynAckTimeout(d time.Duration) ListenerOption {
	return newListenerOption(func(o *listenerOptions) {
		o.synAckTimeout = d
	})
}
func WithListenerHeartTimeout(d time.Duration) ListenerOption {
	return newListenerOption(func(o *listenerOptions) {
		o.heartTimeout = d
	})
}
