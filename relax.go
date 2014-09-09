package gorelax

import (
	"fmt"
	"net/http"
)

//RelaxCompiledPather Comment TODO
type RelaxCompiledPather interface {
	Match(url string) (map[string]string, bool)
}

//RelaxRequester Comment TODO
type RelaxRequester interface {
	Method() string
	Param(key string) string
	Query(key string) string
	Header(key string) string
}

//RelaxResponser Comment TODO
type RelaxResponser interface {
	Send(body string, code int)
}

//RelaxFuncHandler Comment TODO
type RelaxFuncHandler func(req RelaxRequester, res RelaxResponser)

type controller struct {
	compiledPath RelaxCompiledPather
	handler      RelaxFuncHandler
}

//Relax Comment TODO
type Relax struct {
	controllers map[string]controller
}

func (r *Relax) mainHandler(w http.ResponseWriter, req *http.Request) {
	relaxResponse := NewRelaxResponse(&w)

	path := req.Method + req.URL.Path

	for _, controller := range r.controllers {
		if params, ok := controller.compiledPath.Match(path); ok {
			controller.handler(NewRelaxRequest(&params, req), relaxResponse)
			return
		}
	}

	relaxResponse.Send("Not Found", http.StatusNotFound)
}

//RegisterHandler Comment TODO
func (r *Relax) RegisterHandler(method string, url string, funcHandler RelaxFuncHandler) {
	path := method + url
	r.controllers[path] = controller{NewRelaxCompiledPath(path), funcHandler}
}

//Listen Comment TODO
func (r *Relax) Listen(host string, port int) {
	http.HandleFunc("/", r.mainHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}

//NewRelax Comment TODO
func NewRelax() *Relax {
	return &Relax{make(map[string]controller)}
}
