package main

import (
	"log"
	"os"

	"github.com/codecrafters-io/redis-starter-go/internal/server"
)

func main() {
	// parse flags and arguments
	args := os.Args

	// clean the executable arg
	args = args[1:]

	var port string
	var replica server.Replica

	for idx, arg := range args {
		if arg == "--port" {
			port = args[idx+1]
		}

		if arg == "--replicaof" {
			replicaArgs := args[idx+1:]
			if len(replicaArgs) < 2 {
				log.Fatalf("replicaof requires master host and port values respectively")
			}
			replica.MasterHost = replicaArgs[0]
			replica.MasterPort = replicaArgs[1]
		}
	}

	if replica.MasterHost != "" {
		s, err := server.New(port, server.WithReplica(replica))
		if err != nil {
			log.Fatalf("start replica server: %v", err)
		}
		s.Start()
	}

	s, _ := server.New(port)
	s.Start()
}
