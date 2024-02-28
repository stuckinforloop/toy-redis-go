package protocol

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
		return c.set(args)
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

type SetOption string

const (
	EX SetOption = "ex"
	PX SetOption = "px"
)

func (c Command) set(args []string) ([]byte, error) {
	var node cache.Node

	key := args[0]
	node.Value = args[1]

	var option SetOption
	if len(args) > 2 {
		option = SetOption(strings.ToLower(args[2]))
	}

	// NOTE: only handling PX for now
	if option != "" && option == PX {
		ms, err := strconv.Atoi(args[3])
		if err != nil {
			return nil, fmt.Errorf("convert string to int: %w", err)
		}
		node.Expiration = time.Millisecond * time.Duration(ms)
		node.CreatedAt = time.Now()
	}

	if err := cache.Set(key, node); err != nil {
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
