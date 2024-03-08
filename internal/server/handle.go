package server

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/protocol"
)

func (s *Server) handle(conn net.Conn) error {
	defer conn.Close()

	for {
		b := make([]byte, 1024)
		_, err := conn.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			if _, err := conn.Write([]byte("ERR connection read error\r\n")); err != nil {
				return fmt.Errorf("write to conn: %w", err)
			}
		}

		parsedArgs, err := protocol.Parse(b)
		if err != nil {
			return fmt.Errorf("parse request: %w", err)
		}

		args, ok := parsedArgs.([]string)
		if !ok {
			return fmt.Errorf("parse request: %w", err)
		}

		command := strings.ToLower(args[0])

		// TODO: Add support for 2nd element to be a command
		response, err := s.RunCommand(command, args[1:])
		if err != nil {
			return fmt.Errorf("run command: %w", err)
		}

		if _, err := conn.Write(response); err != nil {
			return fmt.Errorf("write to conn: %w", err)
		}
	}

	return nil
}
