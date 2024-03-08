package protocol

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/cache"
)

func Ping() ([]byte, error) {
	msg := fmt.Sprintf("%sPONG\r\n", string(RespString))
	return []byte(msg), nil
}

func Echo(msg string) ([]byte, error) {
	msg = fmt.Sprintf("%s%d\r\n%s\r\n", string(RespBulkString), len(msg), msg)
	return []byte(msg), nil
}

type SetOption string

const (
	EX SetOption = "ex"
	PX SetOption = "px"
)

func Set(args []string) ([]byte, error) {
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

func Get(key string) []byte {
	val, ok := cache.Get(key)
	if !ok {
		return []byte("$-1\r\n")
	}

	msg := fmt.Sprintf("%s%d\r\n%s\r\n", string(RespBulkString), len(val), val)
	return []byte(msg)
}

type Server struct {
	Port    string
	Role    string
	Master  Master
	Replica Replica
}

type Master struct {
	ReplicationID     string
	ReplicationOffset int
}

type Replica struct {
	MasterHost string
	MasterPort string
}

func Info(s Server) []byte {
	// TODO: handle section
	msg := "role" + ":" + s.Role
	if s.Role == "master" {
		msg += "\r\nmaster_replid" + ":" + s.Master.ReplicationID
		msg += "\r\nmaster_repl_offset" + ":" + strconv.Itoa(s.Master.ReplicationOffset)
	}

	msg = fmt.Sprintf("%s%d\r\n%s\r\n", string(RespBulkString), len(msg), msg)
	return []byte(msg)
}
