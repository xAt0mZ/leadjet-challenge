package query

import (
	"encoding/json"
	"fmt"
	"net/http"

	httperror "github.com/xat0mz/leadjet-challenge/api/http/error"
	"github.com/xat0mz/leadjet-challenge/api/http/request"
)

func (h *Handler) queryPost(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {

	var payload queryPayload
	err := request.DecodeJSONPayload(r, &payload)
	if err != nil {
		return &httperror.HandlerError{StatusCode: http.StatusBadRequest, Message: "Invalid request payload", Err: err}
	}

	marshalled, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(marshalled))
	return nil
}
