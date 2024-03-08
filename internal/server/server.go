package server

import (
	"fmt"
	"log"
	"net"
)

const (
	network string = "tcp"

	master string = "master"
	slave  string = "slave"
)

type Server struct {
	port    string
	role    string
	replica Replica
}

type Replica struct {
	MasterHost string
	MasterPort string
}

type Option func(s *Server) error

func WithReplica(r Replica) Option {
	return func(s *Server) error {
		// TODO: validate master is available
		s.replica = r
		s.role = slave

		return nil
	}
}

func New(port string, opts ...Option) (*Server, error) {
	s := &Server{
		port: fmt.Sprintf(":%s", port),
		role: master,
	}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, fmt.Errorf("option: %w", err)
		}
	}

	return s, nil
}

func (s *Server) Start() {
	log.Printf("listening for tcp connections on port %s\n", s.port)

	l, err := net.Listen(network, s.port)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("accept connection: %v", err)
		}

		go func() {
			if err := s.handle(conn); err != nil {
				log.Printf("handle: %v", err)
			}
		}()
	}
}
