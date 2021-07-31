package reverse

import "time"

var defaultDialerOptions = dialerOptions{
	synAck:       true,
	timeout:      time.Second * 75,
	heart:        time.Second * 50,
	heartTimeout: time.Second * 25,
}

type dialerOptions struct {
	synAck       bool
	timeout      time.Duration
	heart        time.Duration
	heartTimeout time.Duration
}
type DialerOption interface {
	apply(*dialerOptions)
}
type funcDialerOption struct {
	f func(*dialerOptions)
}

func (fdo *funcDialerOption) apply(do *dialerOptions) {
	fdo.f(do)
}
func newDialerOption(f func(*dialerOptions)) *funcDialerOption {
	return &funcDialerOption{
		f: f,
	}
}
func WithDialerSynAck(synAck bool) DialerOption {
	return newDialerOption(func(o *dialerOptions) {
		o.synAck = synAck
	})
}
func WithDialerTimeout(timeout time.Duration) DialerOption {
	return newDialerOption(func(o *dialerOptions) {
		o.timeout = timeout
	})
}
func WithDialerHeart(heart time.Duration) DialerOption {
	return newDialerOption(func(o *dialerOptions) {
		o.heart = heart
	})
}
func WithDialerHeartTimeout(timeout time.Duration) DialerOption {
	return newDialerOption(func(o *dialerOptions) {
		o.heartTimeout = timeout
	})
}
