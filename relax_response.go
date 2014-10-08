package gorelax

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//RelaxResponser Comment TODO
type RelaxResponser interface {
	EnableCORS()
	Header(key string, value string)
	Send(body string, code int)
	SendAsJSON(message interface{}, code int)
	Cookie(key string, value string, domain string, path string, age int, secureKey string)
}

//RelaxResponse Comment TODO
type RelaxResponse struct {
	responseWriter *http.ResponseWriter
	req            *http.Request
}

//Send Comment TODO
func (rr *RelaxResponse) Send(body string, code int) {
	response := *(rr.responseWriter)
	response.WriteHeader(code)
	fmt.Fprintf(response, body)
}

//SendAsJSON Comment TODO
func (rr *RelaxResponse) SendAsJSON(message interface{}, code int) {
	result, _ := json.Marshal(message)
	temp := string(result)
	rr.Header("content-type", "application/json; charset=utf-8")
	rr.Send(temp, code)
}

//Header Comment TODO
func (rr *RelaxResponse) Header(key string, value string) {
	response := *(rr.responseWriter)
	response.Header().Set(key, value)
}

//EnableCORS Comment TODO
func (rr *RelaxResponse) EnableCORS() {
	if origin := rr.req.Header.Get("Origin"); origin != "" {
		rr.Header("Access-Control-Allow-Origin", origin)
	}
	rr.Header("Access-Control-Allow-Methods", rr.req.Method)
	rr.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
	rr.Header("Access-Control-Allow-Credentials", "true")
}

//Cookie comment TODO
func (rr *RelaxResponse) Cookie(key string, value string, domain string, path string, age int, secureKey string) {
	cookie := &http.Cookie{
		Name:   key,
		Value:  value,
		Domain: domain,
		Path:   path,
		MaxAge: age,
	}
	response := *(rr.responseWriter)
	http.SetCookie(response, cookie)
}

//NewRelaxResponse Comment TODO
func NewRelaxResponse(responseWriter *http.ResponseWriter, req *http.Request) *RelaxResponse {
	return &RelaxResponse{responseWriter, req}
}
