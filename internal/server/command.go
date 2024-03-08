package server

import "github.com/codecrafters-io/redis-starter-go/internal/protocol"

type Command string

const (
	Ping Command = "ping"
	Echo Command = "echo"
	Get  Command = "get"
	Set  Command = "set"
	Info Command = "info"
)

func (s *Server) RunCommand(command string, args []string) ([]byte, error) {
	switch Command(command) {
	case Ping:
		return protocol.Ping()
	case Echo:
		return protocol.Echo(args[0])
	case Set:
		return protocol.Set(args)
	case Get:
		return protocol.Get(args[0]), nil
	case Info:
		return protocol.Info(s.role), nil
	default:
		return []byte("unkown command"), nil
	}

}
