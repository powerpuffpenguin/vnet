package main

import (
	"flag"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var (
		help, server bool
		addr         string
	)
	flag.BoolVar(&help, `help`, false, `display help`)
	flag.BoolVar(&server, `server`, false, `run server`)
	flag.StringVar(&addr, `addr`, `127.0.0.1:9000`, `run client`)
	flag.Parse()
	if help {
		flag.PrintDefaults()
	} else if server {
		runServer(addr)
	} else {
		runClient(addr)
	}
}
