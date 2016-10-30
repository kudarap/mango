package main

import (
	"github.com/javinc/mango/module"
	"github.com/javinc/mango/test"
)

func main() {
	router := module.Router()

	// Routes consist of a path and a handler function.
	router.Any("/test", test.Handler)
	router.Any("/test/:id", test.Handler)

	router.Run(":8000")
}
