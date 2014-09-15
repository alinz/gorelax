package gorelax

import (
	"fmt"
	"net/http"
)

//Relax Comment TODO
type Relax struct {
	controllers  *RelaxController
	interceptors *RelaxInterceptor
}

func (r *Relax) mainLoop(w http.ResponseWriter, req *http.Request) {
	relaxResponse := NewRelaxResponse(&w, req)

	if controller, params, status := r.controllers.Find(req.Method, req.URL.Path); status == http.StatusOK {
		relaxRequest := NewRelaxRequest(&params, req)
		controller(relaxRequest, relaxResponse)
	} else if status == http.StatusMethodNotAllowed {
		relaxResponse.Send("Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		relaxResponse.Send("Not Found", http.StatusNotFound)
	}
}

//RegisterStaticHandler Comment TODO
func (r *Relax) RegisterStaticHandler(url string, path string) {
	http.Handle(url, http.StripPrefix(url, http.FileServer(http.Dir(path))))
}

//RegisterHandler Comment TODO
func (r *Relax) RegisterHandler(method string, url string, funcHandler RelaxFuncHandler) {
	r.controllers.Add(method, url, funcHandler)
}

//RegisterInterceptorHandler Comment TODO
func (r *Relax) RegisterInterceptorHandler(key string, interceptor RelaxInterceptorHandler) {
	r.interceptors.Add(key, interceptor)
}

//Listen Comment TODO
func (r *Relax) Listen(host string, port int) {
	http.HandleFunc("/", r.mainLoop)
	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}

//NewRelax Comment TODO
func NewRelax() *Relax {
	return &Relax{NewRelaxController(), NewRelaxInterceptor()}
}
