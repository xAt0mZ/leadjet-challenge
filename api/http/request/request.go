// Package request provides function to retrieve content from a *http.Request object of the net/http standard library,.
package request

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONPayload(request *http.Request, v interface{}) error {
	return json.NewDecoder(request.Body).Decode(v)
}
