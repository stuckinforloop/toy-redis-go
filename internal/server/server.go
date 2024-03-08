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

	masterReplicationID string = "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
)

type Server struct {
	port    string
	role    string
	master  Master
	replica Replica
}

type Master struct {
	ReplicationID     string
	ReplicationOffset int
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
	if port == "" {
		port = "6379"
	}
	s := &Server{
		port: fmt.Sprintf(":%s", port),
	}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, fmt.Errorf("option: %w", err)
		}
	}

	if s.role == "" {
		s.role = master
		s.master = Master{
			ReplicationID:     masterReplicationID,
			ReplicationOffset: 0,
		}
	}

	return s, nil
}

func (s *Server) Start() {
	log.Printf("listening for tcp connections on port %s\n", s.port)

	// connect to master server
	if s.role == slave {
		if err := s.connectMaster(); err != nil {
			log.Fatalf("connect to master: %v", err)
		}
	}

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

func (s *Server) connectMaster() error {
	conn, err := net.Dial(network, fmt.Sprintf("%s:%s", s.replica.MasterHost, s.replica.MasterPort))
	if err != nil {
		return fmt.Errorf("connect to master: %w", err)
	}

	if _, err := conn.Write([]byte("*1\r\n$4\r\nping\r\n")); err != nil {
		return fmt.Errorf("ping master: %w", err)
	}

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("read master response: %w", err)
	}

	log.Println("resp from master: ", string(buf[:n]))

	return nil
}
