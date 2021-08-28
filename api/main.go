// main.go
package main

import (
	"log"

	"github.com/xat0mz/leadjet-challenge/api/http"
)

func buildServer() http.IServer {
	return &http.Server{
		Port: 10000,
	}
}

func main() {
	server := buildServer()
	server.Init()
	log.Printf("[INFO] [main] Starting server\n")
	err := server.Start()
	log.Printf("[INFO] [main] Http server exited: %s\n", err)
}
