package handler

import (
	"net/http"
	"strings"

	"github.com/xat0mz/leadjet-challenge/api/query"
)

// Handler is the HTTP handler used to handle registry operations.
type Handler struct {
	QueryHandler *query.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/api/query"):
		http.StripPrefix("/api", h.QueryHandler).ServeHTTP(w, r)
	}
}
