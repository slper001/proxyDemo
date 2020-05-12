package proxy

import (
	"io"
	"net"
)

type Proxy struct {
	listener       net.Listener
	backendAddress string
}

func NewProxy(agentAddress, serverAddress string) *Proxy {
	lst, _ := net.Listen("tcp", agentAddress)
	return &Proxy{
		listener:       lst,
		backendAddress: serverAddress,
	}
}

func (p *Proxy) Run() {
	for {
		frontConn, _ := p.listener.Accept()
		backendConn, _ := net.Dial("tcp", p.backendAddress)
		go p.Agent(frontConn, backendConn)
	}
}

func (p *Proxy) Agent(frontConn, backendConn net.Conn) {
	go io.Copy(backendConn, frontConn)
	go io.Copy(frontConn, backendConn)
}
