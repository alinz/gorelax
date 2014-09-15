package main

import (
	"fmt"
	relax "github.com/alinz/gorelax"
)

func main() {
	relax := gorelax.NewRelax()
	relax.RegisterHandler("GET", "/user/{id}", func(req gorelax.RelaxRequester, res gorelax.RelaxResponser) {
		res.EnableCORS()
		res.Cookie("user", "hello", "", "", -1, "")
		res.Send(fmt.Sprintf("your cookie is %s", req.Cookie("user")), 200)
	})

	relaxServer.Listen(9000)
}
