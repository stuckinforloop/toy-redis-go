package server

import (
	"fmt"
	"log"
	"net"
)

const (
	network string = "tcp"
)

type Server struct {
	addr    string
	network string
}

func New(port *string) *Server {
	return &Server{
		addr:    fmt.Sprintf(":%s", *port),
		network: network,
	}
}

func (s *Server) Start() {
	log.Printf("listening for tcp connections on port %s\n", s.addr)

	l, err := net.Listen(s.network, s.addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("accept connection: %v", err)
		}

		go func() {
			if err := handle(conn); err != nil {
				log.Printf("handle: %v", err)
			}
		}()
	}
}
