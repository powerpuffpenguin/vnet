package vnet

import (
	"context"
	"net"
)

type Dialer interface {
	Dial(network, addr string) (net.Conn, error)
	DialContext(ctx context.Context, network, addr string) (conn net.Conn, e error)
}
