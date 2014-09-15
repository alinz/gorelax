package gorelax

import (
	"github.com/alinz/gorelax/util"
)

//RelaxNextInteceotorHandler Comment TODO
type RelaxNextInteceotorHandler func()

//RelaxInterceptorHandler Comment TODO
type RelaxInterceptorHandler func(req *RelaxRequester, res *RelaxResponser, next RelaxNextInteceotorHandler)

//RelaxInterceptor Comment TODO
type RelaxInterceptor struct {
	interceptors util.MapArray
}

type internalRelaxInteceptor struct {
	interceptor RelaxInterceptorHandler
}

//Add Comment TODO
func (ri *RelaxInterceptor) Add(key string, interceptor RelaxInterceptorHandler) {
	ri.interceptors.Add(key, interceptor)
}

//Process Comment TODO
func (ri *RelaxInterceptor) Process(key string, req *RelaxRequester, res *RelaxResponser) {

	var next RelaxNextInteceotorHandler

	if interceptorHandlers, ok := ri.interceptors.Get(key); ok {
		index := 0
		length := len(interceptorHandlers)
		next = func() {
			if index < length {
				interceptorHandlers[index].(internalRelaxInteceptor).interceptor(req, res, next)
				index++
			}
		}

		next()
	}
}

//NewRelaxInterceptor Comment TODO
func NewRelaxInterceptor() *RelaxInterceptor {
	return &RelaxInterceptor{util.MakeMapArray()}
}
