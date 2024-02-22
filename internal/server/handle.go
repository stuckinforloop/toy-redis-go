package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/redis-starter-go/internal/protocol"
)

func handle(conn net.Conn) error {
	defer conn.Close()

	// TODO parse request data

	// response for ping command
	firstByte := protocol.DataTypeToFirstByte[protocol.SimpleString]
	response := firstByte + "PONG" + "\r\n"

	if _, err := conn.Write([]byte(response)); err != nil {
		return fmt.Errorf("write to conn: %w", err)
	}

	return nil
}
