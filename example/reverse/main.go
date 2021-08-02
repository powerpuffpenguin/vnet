package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/powerpuffpenguin/vnet"
	"github.com/powerpuffpenguin/vnet/reverse"
)

func main() {
	l, e := net.Listen(`tcp`, `127.0.0.1:9000`)
	if e != nil {
		log.Fatalln(e)
	}

	var (
		dialer *reverse.Dialer = reverse.NewDialer(l)
		d      vnet.Dialer     = dialer
	)
	go dialer.Serve()
	l = reverse.Listen(l.Addr())

	// run client
	go runClient(&http.Client{
		Transport: &http.Transport{
			// http client by vnet.Dialer
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				fmt.Println(network, addr)
				return d.DialContext(ctx, network, addr)
			},
		},
	})

	// run server
	runServer(l)
}

func runClient(client *http.Client) {
	// get /info
	resp, e := client.Get(`http://reverse/info`)
	if e != nil {
		log.Fatalln(e)
	}
	b, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Printf("/info resp: %s\n", b)

	// get /exit
	resp, e = client.Get(`http://reverse/exit`)
	if e != nil {
		log.Fatalln(e)
	}
	b, e = ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Printf("/exit resp: %s\n", b)
}

func runServer(l net.Listener) {
	mux := http.NewServeMux()

	mux.HandleFunc(`/info`, func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`reverse listener`))
	})
	mux.HandleFunc(`/exit`, func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`success`))
		go func() {
			time.Sleep(time.Millisecond * 100)
			l.Close()
		}()
	})

	e := http.Serve(l, mux)
	if e != nil && !errors.Is(e, vnet.ErrListenerClosed) {
		log.Fatalln(e)
	}
}
