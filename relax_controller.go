package gorelax

import (
	"net/http"

	"github.com/alinz/gorelax/util"
)

//RelaxFuncHandler Comment TODO
type RelaxFuncHandler func(req RelaxRequester, res RelaxResponser)

type controller struct {
	compiledPath RelaxCompiledPather
	handler      RelaxFuncHandler
}

//RelaxController Comment TODO
type RelaxController struct {
	controllers util.MapArray
}

//Add Comment TODO
func (c *RelaxController) Add(method string, url string, funcHandler RelaxFuncHandler) {
	c.controllers.Add(method, controller{NewRelaxCompiledPath(url), funcHandler})
}

//Find Comment TODO
func (c *RelaxController) Find(method string, url string) (RelaxFuncHandler, map[string]string, int) {
	var params map[string]string
	var handler RelaxFuncHandler

	status := http.StatusMethodNotAllowed

	c.controllers.Iterate(method, func(ctrl interface{}, index int) bool {
		if paramValue, ok := ctrl.(controller).compiledPath.Match(url); ok {
			params = paramValue
			handler = ctrl.(controller).handler
			status = http.StatusOK
			return false
		}
		status = http.StatusNotFound
		return true
	})

	return handler, params, status
}

//NewRelaxController Comment TODO
func NewRelaxController() *RelaxController {
	return &RelaxController{util.MakeMapArray()}
}
