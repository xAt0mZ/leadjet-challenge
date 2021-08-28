package handler

import (
	"net/http"
	"strings"

	"github.com/xat0mz/leadjet-challenge/api/queries"
)

// Handler is the HTTP handler used to handle registry operations.
type Handler struct {
	QueriesHandler *queries.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/api/queries"):
		http.StripPrefix("/api", h.QueriesHandler).ServeHTTP(w, r)
	}
}
