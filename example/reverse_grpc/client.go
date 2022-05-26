package main

import (
	"context"
	"log"
	"net"
	grpc_math "reverse_grpc/math"
	"time"

	"github.com/powerpuffpenguin/vnet"
	"github.com/powerpuffpenguin/vnet/reverse"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runClient(addr string) {
	l, e := net.Listen(`tcp`, addr)
	if e != nil {
		log.Fatalln(e)
	}
	dialer := reverse.NewDialer(l)
	cc, e := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(c context.Context, s string) (net.Conn, error) {
			return dialer.DialContext(c, `tcp`, s)
		}),
	)
	if e != nil {
		log.Fatalln(e)
	}
	defer cc.Close()

	go func() {
		for {
			time.Sleep(time.Second)
			// request server service
			runRequest(cc)
		}
	}()

	e = dialer.Serve()
	if e != vnet.ErrDialerClosed {
		log.Fatalln(e)
	}
}
func runRequest(cc *grpc.ClientConn) {
	client := grpc_math.NewMathClient(cc)
	vals := []int64{1, 2, 3}
	log.Println(`add`, vals)
	resp, e := client.Add(context.Background(), &grpc_math.AddRequest{
		Vals: vals,
	})
	if e != nil {
		log.Println(e)
		return
	}
	log.Println(`sum`, resp.Sum)

	stream, e := client.Random(context.Background())
	if e != nil {
		log.Println(e)
		return
	}
	for i := 0; i < 3; i++ {
		max := int64(100 * (i + 1))
		log.Println(`randon`, max)
		e = stream.Send(&grpc_math.RandomRequest{
			Max: max,
		})
		if e != nil {
			log.Println(e)
			return
		}
		resp, e := stream.Recv()
		if e != nil {
			log.Println(e)
			return
		}
		log.Println(`val`, resp.Val)
	}
	stream.CloseSend()
}
