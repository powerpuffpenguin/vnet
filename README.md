# vnet
virtual net interface

golang provides net.Listener and net.Conn interfaces to wrap the network. A lot of interesting tcp codes are built on these two interfaces. So as long as these two interfaces are met, these codes can be used outside of tcp, this is what vnet does.

# PipeListener

The ListenPipe function returns a PipeListener. PipeListener implements the net.Listener interface in memory. Use PipeListener.Dial or PipeListener.DialContext to dial.

One of the most typical applications is to use PipeListener to provide grpc and grpc-gateway services in the program. The grpc and grpc-gateway use memory copy communication instead of socket communication.

```
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
```

# reverse.Dialer reverse.Listener

reverse.Dialer reverse.Listener used to reverse the server client. This allows socket.connect to work in the server role and socket.accept to work in the client role.

The most typical use is that the server role works on the internal network, and the client role works on the external network. For example, reverse Trojan horses or programs that need to break through the intranet to provide services. Imagine how easy it is to write a reverse Trojan using grpc.

```
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
```