package main

import (
	"flag"
	"log"

	"github.com/codecrafters-io/redis-starter-go/internal/server"
)

func main() {
	// flag variables
	var port string
	var replicaOf string

	flag.StringVar(&port, "port", "6379", "port on which redis server will listen for requests")
	flag.StringVar(&replicaOf, "replicaof", "", "replica server details")

	flag.Parse()

	if replicaOf != "" {
		// TODO: parse replica flag arguments

		replica := server.Replica{
			MasterHost: "",
			MasterPort: "",
		}

		s, err := server.New(port, server.WithReplica(replica))
		if err != nil {
			log.Fatalf("create new replica server: %v", err)
		}

		s.Start()
	}

	s, _ := server.New(port)
	s.Start()
}
