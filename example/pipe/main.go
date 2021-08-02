package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"

	"github.com/powerpuffpenguin/vnet"
)

func main() {
	var (
		// listen pipe
		p = vnet.ListenPipe()

		l net.Listener = p
		d vnet.Dialer  = p
	)

	// run client
	go runClient(&http.Client{
		Transport: &http.Transport{
			// http client by vnet.Dialer
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return d.DialContext(ctx, network, addr)
			},
		},
	})

	// run server
	runServer(l)
}

func runClient(client *http.Client) {
	// get /info
	resp, e := client.Get(`http://pipe/info`)
	if e != nil {
		log.Fatalln(e)
	}
	b, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Printf("/info resp: %s\n", b)

	// get /exit
	resp, e = client.Get(`http://pipe/exit`)
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
		rw.Write([]byte(`pipe listener`))
	})
	mux.HandleFunc(`/exit`, func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`success`))
		runtime.Gosched()
		l.Close()
	})

	e := http.Serve(l, mux)
	if e != nil && !errors.Is(e, vnet.ErrListenerClosed) {
		log.Fatalln(e)
	}
}
