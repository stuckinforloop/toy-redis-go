package main

import (
	"flag"

	"github.com/codecrafters-io/redis-starter-go/internal/server"
)

func main() {
	port := flag.String("port", "6379", "port on which redis server will listen for requests")
	flag.Parse()

	server.New(port).Start()
}
