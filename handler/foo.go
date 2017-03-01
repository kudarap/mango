package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/javinc/mango/model"
	"github.com/javinc/mango/service"
	"github.com/javinc/mango/service/logic"
	"github.com/javinc/mango/store"
)

// FooHandler http call
func FooHandler(c *gin.Context) {
	// mocking user detail on request
	service.ToContext(c, logic.New(model.User{
		ID:   "testid",
		Type: "client",
	}))

	s, _ := service.GetFoo(c, "testid")

	c.String(http.StatusOK, "%s", s)
}

// FooHandlerx http call
func FooHandlerx(c *gin.Context) {
	s, _ := store.FromContext(c).GetFoo("testid")

	c.String(http.StatusOK, "%s", s)
}
