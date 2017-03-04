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
	r.POST("/register", user.RegisterHandler)
	r.POST("/login", user.LoginHandler)

	// read only
	r.Static("/upload", "./upload")

	// private group
	auth := r.Group("/", module.AuthRequired())
	{
		auth.Any("/test", test.Handler)
		auth.Any("/test/:id", test.Handler)

		auth.GET("/me", user.MeHandler)
		auth.Any("/user", user.Handler)
		auth.Any("/user/:id", user.Handler)

		auth.Any("/file", file.Handler)
		auth.Any("/file/:id", file.Handler)
		auth.POST("/upload", file.UploadHandler)
	}

	r.Run(module.Config.Host)

	// graceful shutdown
	// endless.ListenAndServe(module.Config.Host, r)
}
