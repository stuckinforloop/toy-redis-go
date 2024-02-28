package protocol

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/internal/cache"
)

type Command string

const (
	Ping Command = "ping"
	Echo Command = "echo"
	Get  Command = "get"
	Set  Command = "set"
)

func (c Command) Run(args []string) ([]byte, error) {
	switch c {
	case Ping:
		return c.ping()
	case Echo:
		return c.echo(args[0])
	case Set:
		return c.set(args[0], args[1])
	case Get:
		return c.get(args[0]), nil
	default:
		return []byte("unkown command"), nil
	}
}

func (c Command) ping() ([]byte, error) {
	msg := fmt.Sprintf("%sPONG\r\n", string(RespString))
	return []byte(msg), nil
}

func (c Command) echo(msg string) ([]byte, error) {
	msg = fmt.Sprintf("%s%d\r\n%s\r\n", string(RespBulkString), len(msg), msg)
	return []byte(msg), nil
}

func (c Command) set(key string, val any) ([]byte, error) {
	if err := cache.Set(key, val); err != nil {
		return nil, fmt.Errorf("set value: %w", err)
	}

	return []byte("+OK\r\n"), nil
}

func (c Command) get(key string) []byte {
	val, ok := cache.Get(key)
	if !ok {
		return []byte("$-1\r\n")
	}

	msg := fmt.Sprintf("%s%d\r\n%s\r\n", string(RespBulkString), len(val), val)
	return []byte(msg)
}
