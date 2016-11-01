package main

import (
	"github.com/javinc/mango/module"
	"github.com/javinc/mango/test"
	"github.com/javinc/mango/user"
)

func main() {
	r := module.Router()

	// Routes consist of a path and a handler function.
	r.Any("/test", test.Handler)
	r.Any("/test/:id", test.Handler)

	r.Any("/login", user.LoginHandler)
	r.Any("/me", user.MeHandler)

	r.Any("/user", user.Handler)
	r.Any("/user/:id", user.Handler)

	r.Run(module.Config.Host)

	// graceful shutdown
	// endless.ListenAndServe(module.Config.Host, r)
}
