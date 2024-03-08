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
		server := protocol.Server{
			Port: s.port,
			Role: s.role,
		}

		switch s.role {
		case master:
			server.Master.ReplicationID = s.master.ReplicationID
			server.Master.ReplicationOffset = s.master.ReplicationOffset
		case slave:
			server.Replica.MasterHost = s.replica.MasterHost
			server.Replica.MasterPort = s.replica.MasterPort
		}

		return protocol.Info(server), nil
	default:
		return []byte("unkown command"), nil
	}

}
