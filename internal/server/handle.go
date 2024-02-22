package server

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/codecrafters-io/redis-starter-go/internal/protocol"
)

func handle(conn net.Conn) error {
	defer conn.Close()

	b := make([]byte, 1024)
	for {
		_, err := conn.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("connection read: %w", err)
		}

		// TODO parse request data

		firstByte := protocol.DataTypeToFirstByte[protocol.SimpleString]
		response := firstByte + "PONG" + "\r\n"

		if _, err := conn.Write([]byte(response)); err != nil {
			return fmt.Errorf("write to conn: %w", err)
		}
	}

	return nil
}
