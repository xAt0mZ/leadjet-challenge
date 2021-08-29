package query

import (
	"net/http"

	"github.com/gorilla/mux"
	httperror "github.com/xat0mz/leadjet-challenge/api/http/error"
)

// Handler is the HTTP handler used to handle queries operations.
type Handler struct {
	*mux.Router
}

// NewHandler creates a handler to manage queries operations.
func NewHandler() *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}
	h.initRouter()

	return h
}

// initRouter binds paths and methods to handler functions
func (h *Handler) initRouter() {
	h.Handle("/query", httperror.LoggerHandler(h.queryPost)).Methods(http.MethodPost)
}
