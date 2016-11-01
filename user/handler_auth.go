package user

import (
	"github.com/gin-gonic/gin"
	x "github.com/javinc/mango/module"
)

// RegisterHandler user register
func RegisterHandler(c *gin.Context) {
	switch c.Request.Method {
	case x.POST:
		var payload Object

		err := c.BindJSON(&payload)
		if err != nil {
			x.Panic("JSON_BIND_ERROR", err.Error())

			return
		}

		user, err := service.Register(payload)
		if err != nil {
			x.Error("USER_EMAIL_REGISTER_ERROR", err.Error())

			return
		}

		// success
		x.Output(user)

		return
	}

	x.MethodNotAllowedError()
}

// LoginHandler user login
func LoginHandler(c *gin.Context) {
	type Login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	switch c.Request.Method {
	case x.POST:
		var payload Login

		err := c.BindJSON(&payload)
		if err != nil {
			x.Panic("USER_JSON_BIND_ERROR", err.Error())

			return
		}

		auth, err := service.Login(payload.Email, payload.Password)
		if err != nil {
			x.Error("USER_EMAIL_LOGIN_ERROR", err.Error())

			return
		}

		if err == nil {
			x.Output(auth)

			return
		}

		x.Error("INVALID_AUTH", "unauthorize user")

		return
	}

	x.MethodNotAllowedError()
}

// MeHandler check authentication
func MeHandler(c *gin.Context) {
	switch c.Request.Method {
	case x.GET:
		user, err := service.Get(x.GetAuthUser().ID)
		if err != nil {
			x.Error("INVALID_USER", "no user found with this token")

			return
		}

		x.Output(user)

		return
	}

	x.MethodNotAllowedError()
}
