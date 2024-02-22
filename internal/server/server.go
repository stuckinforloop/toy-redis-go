package server

import (
	"log"
	"net"
)

const (
	addr    string = ":6379"
	network string = "tcp"
)

type Server struct {
	addr    string
	network string
}

func New() *Server {
	return &Server{
		addr:    addr,
		network: network,
	}
}

func (s *Server) Start() {
	log.Printf("listening for tcp connections on port: %s\n", s.addr)

	l, err := net.Listen(s.network, s.addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	for {
		_, err := l.Accept()
		if err != nil {
			log.Fatalf("accept connection: %v", err)
		}
	}
}
