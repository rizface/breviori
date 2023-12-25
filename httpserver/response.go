package httpserver

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func writeHTTPResponse(w http.ResponseWriter, resp httpResponse) {
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}
