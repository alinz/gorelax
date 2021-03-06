package gorelax

import (
	"net/http"
	"net/url"
)

//RelaxRequester Comment TODO
type RelaxRequester interface {
	Method() string
	Param(key string) string
	Query(key string) string
	Header(key string) string
	Cookie(key string) string
}

//RelaxRequest Comment TODO
type RelaxRequest struct {
	params  *map[string]string
	req     *http.Request
	queries url.Values
}

//Method Comment TODO
func (rr *RelaxRequest) Method() string {
	return rr.req.Method
}

//Param Comment TODO
func (rr *RelaxRequest) Param(key string) string {
	mp := *rr.params
	return mp[key]
}

//Query Comment TODO
func (rr *RelaxRequest) Query(key string) string {
	if rr.queries == nil {
		rr.queries = rr.req.URL.Query()
	}
	return rr.queries.Get(key)
}

//Header Comment TODO
func (rr *RelaxRequest) Header(key string) string {
	return rr.req.Header.Get(key)
}

//Cookie Comment TODO
func (rr *RelaxRequest) Cookie(key string) string {
	cookie, err := rr.req.Cookie(key)
	if err == nil {
		return cookie.Value
	}
	return ""
}

//NewRelaxRequest Comment TODO
func NewRelaxRequest(params *map[string]string, req *http.Request) *RelaxRequest {
	return &RelaxRequest{params, req, nil}
}
