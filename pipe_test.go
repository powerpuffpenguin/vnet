package vnet_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"testing"

	"github.com/powerpuffpenguin/vnet"
)

func TestPipe(t *testing.T) {
	var (
		// listen pipe
		p = vnet.ListenPipe()

		l net.Listener = p
		d vnet.Dialer  = p
	)

	// run client
	ch := make(chan error, 2)
	go runClient(
		&http.Client{
			Transport: &http.Transport{
				// http client by vnet.Dialer
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return d.DialContext(ctx, network, addr)
				},
			},
		},
		ch,
	)

	// run server
	go runServer(l, ch)

	for i := 0; i < 2; i++ {
		e := <-ch
		if e != nil {
			t.Fatal(e)
		}
	}
}
func runClient(client *http.Client, ch chan<- error) {
	// get /info
	resp, e := client.Get(`http://pipe/info`)
	if e != nil {
		ch <- e
		return
	}
	b, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		ch <- e
		return
	}
	fmt.Printf("/info resp: %s\n", b)

	// get /exit
	resp, e = client.Get(`http://pipe/exit`)
	if e != nil {
		ch <- e
		return
	}
	b, e = ioutil.ReadAll(resp.Body)
	if e != nil {
		ch <- e
		return
	}
	fmt.Printf("/exit resp: %s\n", b)
	ch <- nil
}

func runServer(l net.Listener, ch chan<- error) {
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
		ch <- e
	} else {
		ch <- nil
	}
}
