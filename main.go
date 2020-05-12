package main

import (
	"github.com/proxyDemo/client"
	"github.com/proxyDemo/proxy"
	"github.com/proxyDemo/server"
)

func main() {
	servers := server.NewServer(":9998", ":9999")
	for _, s := range servers {
		go s.Run()
	}

	p := proxy.NewProxy(":9000", ":9999")
	go p.Run()

	for _, c := range client.NewClient(":9000", 3) {
		go c.Run()
	}
	select {}
}
