package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/javinc/mango/config"
	foo "github.com/javinc/mango/foo/controller"
	"github.com/javinc/mango/rest"
)

func main() {
	log.Println("[main]", "starting...")

	c := config.Get()
	r := rest.Router()

	// routing
	route(r)

	log.Println("[main]", "listening and serving on", c.Port)
	r.Run(c.Port)
}

func route(r *gin.Engine) {
	r.Any("/", foo.Handler)
}
