package main

import (
	"github.com/javinc/mango/file"
	"github.com/javinc/mango/module"
	"github.com/javinc/mango/test"
	"github.com/javinc/mango/user"
)

func main() {
	r := module.Router()

	// public endpoint
	r.Any("/login", user.LoginHandler)
	r.Any("/me", user.MeHandler)

	r.Any("/test", test.Handler)
	r.Any("/test/:id", test.Handler)

	r.Any("/user", user.Handler)
	r.Any("/user/:id", user.Handler)

	r.Any("/file", file.Handler)
	r.Any("/file/:id", file.Handler)

	r.Run(module.Config.Host)

	// graceful shutdown
	// endless.ListenAndServe(module.Config.Host, r)
}
