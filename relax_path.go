package gorelax

import (
	"regexp"
	"strings"
)

var (
	lookFor     = []string{"{", "}"}
	replaceWith = []string{"(?P<", ">[0-9a-zA-Z_]+)"}
)

//RelaxCompiledPather Comment TODO
type RelaxCompiledPather interface {
	Match(url string) (map[string]string, bool)
}

func compiledURL(path string) *regexp.Regexp {
	for index := range lookFor {
		path = strings.Replace(path, lookFor[index], replaceWith[index], -1)
	}
	path = "^" + path + "$"
	r, _ := regexp.Compile(path)

	return r
}

//RelaxCompiledPath Comment TODO
type RelaxCompiledPath struct {
	compiledPath *regexp.Regexp
}

//Match Comment TODO
func (rcp *RelaxCompiledPath) Match(path string) (map[string]string, bool) {
	var params map[string]string
	result := false

	if rcp.compiledPath.MatchString(path) {
		params = rcp.extractParams(path)
		result = true
	}

	return params, result
}

func (rcp *RelaxCompiledPath) extractParams(target string) map[string]string {
	result := make(map[string]string)
	keys := rcp.compiledPath.SubexpNames()
	values := rcp.compiledPath.FindStringSubmatch(target)

	for index, value := range keys {
		if index != 0 {
			result[value] = values[index]
		}
	}

	return result
}

//NewRelaxCompiledPath Comment TODO
func NewRelaxCompiledPath(path string) *RelaxCompiledPath {
	return &RelaxCompiledPath{compiledURL(path)}
}
