// Package error provides error/logging functions that can be used in conjuction with http.Handler.
package httperror

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	// LoggerHandler defines a HTTP handler that includes a HandlerError return pointer
	LoggerHandler func(http.ResponseWriter, *http.Request) *HandlerError

	// HandlerError represents an error raised inside a HTTP handler
	HandlerError struct {
		StatusCode int
		Message    string
		Err        error
	}

	errorResponse struct {
		Message string `json:"message,omitempty"`
		Details string `json:"details,omitempty"`
	}
)

// LoggerHandler implements net/http/Handler interface
func (handler LoggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := handler(w, r)
	if err != nil {
		writeErrorResponse(w, err)
	}
}

func writeErrorResponse(w http.ResponseWriter, err *HandlerError) {
	log.Printf("http error: %s (err=%s) (code=%d)\n", err.Message, err.Err, err.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(&errorResponse{Message: err.Message, Details: err.Err.Error()})
}

// WriteError is a convenience function that creates a new HandlerError before calling writeErrorResponse.
// For use outside of the standard http handlers.
func WriteError(w http.ResponseWriter, code int, message string, err error) {
	writeErrorResponse(w, &HandlerError{code, message, err})
}
