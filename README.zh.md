# vnet
虛擬網路接口

[English](README.md)

golang 提供了 net.Listener 和 net.Conn 接口來封裝網絡。 在這兩個接口上構建了很多有趣的 tcp 代碼。 所以只要滿足這兩個接口，這些代碼就可以在tcp之外使用，這就是 vnet 要做的事。

* [PipeListener](#pipelistener)
* [reverse.Dialer reverse.Listener](#reversedialer-reverselistener)

# PipeListener

ListenPipe 函數返回一個 PipeListener。 PipeListener 在內存中實現 net.Listener 接口。 使用 PipeListener.Dial 或 PipeListener.DialContext 進行撥號。

最典型的應用之一就是在程序中使用 PipeListener 提供 grpc 和 grpc-gateway 服務。 grpc 和 grpc-gateway 使用內存複製通信而不是套接字通信。

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

reverse.Dialer reverse.Listener 用於逆反服務器客戶端。 這允許 socket.connect 以服務器角色工作，而 socket.accept 以客戶端角色工作。

最典型的使用場景是服務器工作在內網，客戶端工作在外網所以需要服務器撥號到客戶端。 比如逆向木馬或者需要突破內網提供服務的程序。 想像一下，使用 grpc 編寫反向木馬是多麼容易。

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