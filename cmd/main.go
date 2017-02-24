package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/javinc/mango/config"
	"github.com/javinc/mango/handler"
	"github.com/javinc/mango/model"
	"github.com/javinc/mango/rest"
	"github.com/javinc/mango/store"
	"github.com/javinc/mango/store/resource"
)

func main() {
	log.Println("[main]", "starting...")

	c := config.Get()
	r := rest.Router()

	r.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			log.Println("[main]", "middleware...")

			// mocking user detail on request
			store.ToContext(c, resource.New(model.User{
				ID:   "testid",
				Type: "client",
			}))

			c.Next()
		}
	}())

	// routing
	route(r)

	log.Println("[main]", "listening and serving on", c.Port)
	r.Run(c.Port)
}

func route(r *gin.Engine) {
	r.Any("/", handler.FooHandler)
}
