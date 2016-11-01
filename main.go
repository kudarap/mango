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
	r.Any("/register", user.RegisterHandler)
	r.Any("/login", user.LoginHandler)

	// Authorization group
	auth := r.Group("/", module.AuthRequired())
	{
		auth.Any("/me", user.MeHandler)

		auth.Any("/test", test.Handler)
		auth.Any("/test/:id", test.Handler)

		auth.Any("/user", user.Handler)
		auth.Any("/user/:id", user.Handler)

		auth.Any("/file", file.Handler)
		auth.Any("/file/:id", file.Handler)
	}

	r.Run(module.Config.Host)

	// graceful shutdown
	// endless.ListenAndServe(module.Config.Host, r)
}
