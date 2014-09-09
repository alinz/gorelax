package gorelax

import (
	"fmt"
	"net/http"
)

//RelaxResponse Comment TODO
type RelaxResponse struct {
	responseWriter *http.ResponseWriter
}

//Send Comment TODO
func (rr *RelaxResponse) Send(body string, code int) {
	response := *(rr.responseWriter)
	response.WriteHeader(code)
	fmt.Fprintf(response, body)
}

//NewRelaxResponse Comment TODO
func NewRelaxResponse(responseWriter *http.ResponseWriter) *RelaxResponse {
	return &RelaxResponse{responseWriter}
}
