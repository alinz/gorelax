## GoRelax is a Micro Framework to write RESTful API in Go.

### Features
1. Supports parameters extraction
2. Supports basic params checks whether a value is a number or string.
3. Simple and abstract http requests

### Usage

create a Relax struct and register all your handlers by using `Relax.RegisterHandler`. Then call `Relax.Listen` to start listining on your port.

```go
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
```

For full example, please take a look at the example folder.
