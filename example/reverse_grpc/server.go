package main

import (
	"io"
	"log"

	grpc_math "reverse_grpc/math"
	"reverse_grpc/server"

	"github.com/powerpuffpenguin/vnet/reverse"
	"google.golang.org/grpc"
)

type Addr string

func (a Addr) String() string {
	return string(a)
}
func (a Addr) Network() string {
	return `tcp`
}
func runServer(addr string) {
	l := reverse.Listen(Addr(addr))

	s := grpc.NewServer()

	grpc_math.RegisterMathServer(s, server.Server{})

	e := s.Serve(l)
	if e != io.EOF {
		log.Fatalln(e)
	}
}
