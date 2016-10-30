package main

import (
	"github.com/javinc/mango/module"
	"github.com/javinc/mango/test"
	"github.com/javinc/mango/user"
)

func main() {
	router := module.Router()

	// Routes consist of a path and a handler function.
	router.Any("/test", test.Handler)
	router.Any("/test/:id", test.Handler)

	router.Any("/user", user.Handler)
	router.Any("/user/:id", user.Handler)

	router.Run(":8000")
}
