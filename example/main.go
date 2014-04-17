package main

import (
	"fmt"
	relax "github.com/alinz/gorelax"
)

func main() {
	var relaxServer relax.Relax = relax.Relax{}

	relaxServer.RegisterHandler("GET", "/user/{id}<number>", func(req relax.RelaxHttpRequest, res relax.RelaxHttpResponse) {
		res.Send(fmt.Sprintf("You have Selected user %s", req.Params["id"]), 200)
	})

	relaxServer.RegisterHandler("GET", "/users", func(req relax.RelaxHttpRequest, res relax.RelaxHttpResponse) {
		res.Send("[]", 200)
	})

	relaxServer.Listen(9000)
}
