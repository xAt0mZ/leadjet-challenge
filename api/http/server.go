package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xat0mz/leadjet-challenge/api/http/handler"
	"github.com/xat0mz/leadjet-challenge/api/queries"
)

type IServer interface {
	Init()
	Start() error
}

// Server implements IServer interface
type Server struct {
	Port int

	handler *handler.Handler
}

func (server *Server) Init() {
	server.handler = &handler.Handler{
		QueriesHandler: queries.NewHandler(),
	}
}

func (server *Server) Start() error {
	log.Printf("[INFO] [http,server] [message: starting HTTP server on port %d]", server.Port)

	httpServer := &http.Server{
		Addr:    ":" + fmt.Sprint(server.Port),
		Handler: server.handler,
	}
	return httpServer.ListenAndServe()
}
