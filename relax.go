package relax

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var look_for = []string{"{", "}", "<number>", "<string>"}
var replace_with = []string{"(?P<", ">", "[0-9\\.]+)", "[0-9a-zA-Z_]+)"}

func compileActualUrl(actualUrl string) (*regexp.Regexp, string) {
	for index, _ := range look_for {
		actualUrl = strings.Replace(actualUrl, look_for[index], replace_with[index], -1)
	}
	actualUrl = "^" + actualUrl + "$"
	r, _ := regexp.Compile(actualUrl)

	return r, actualUrl
}

func applyPath(target string, source *regexp.Regexp) map[string]string {
	result := make(map[string]string)
	if source.MatchString(target) {
		keys := source.SubexpNames()
		values := source.FindStringSubmatch(target)

		for index, value := range keys {
			if index != 0 {
				result[value] = values[index]
			}
		}
	}
	return result
}

type RelaxHttpRequest struct {
	Params  map[string]string
	request *http.Request
}

type RelaxHttpResponse struct {
	responseWriter http.ResponseWriter
}

func (self RelaxHttpResponse) Send(body string, code int) {
	self.responseWriter.WriteHeader(code)
	fmt.Fprintf(self.responseWriter, body)
}

type RelaxHttpFuncHandler func(req RelaxHttpRequest, res RelaxHttpResponse)

type controller struct {
	compiledPath *regexp.Regexp
	method       string
	httpHandler  RelaxHttpFuncHandler
}

var controllers = make(map[string]*controller)

type Relax struct{}

func (self Relax) mainHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL

	method := r.Method

	var targetController *controller = nil

	for _, value := range controllers {
		if value.compiledPath.MatchString(url.Path) {
			if value.method == method {
				targetController = value
				break
			}
		}
	}

	if targetController != nil {
		params := applyPath(url.Path, targetController.compiledPath)
		targetController.httpHandler(RelaxHttpRequest{params, r}, RelaxHttpResponse{w})
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found")
	}
}

func (self Relax) RegisterHandler(method string, path string, httpFuncHandler RelaxHttpFuncHandler) bool {
	regexpPath, actualPath := compileActualUrl(path)
	if value, ok := controllers["foo"]; ok && value.method == method {
		return false
	} else {
		controllers[actualPath] = &controller{regexpPath, method, httpFuncHandler}
		return true
	}
}

func (self Relax) Listen(port int) {
	http.HandleFunc("/", self.mainHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
