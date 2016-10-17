package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login test
type Login struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

var service Service

// Handler test
func Handler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:

		d := service.Find()

		c.JSON(http.StatusOK, d)

		return
	case http.MethodPost:
		c.JSON(http.StatusOK, gin.H{
			"msg": "welcome POST",
		})

		return
	case http.MethodPut:
		c.JSON(http.StatusOK, gin.H{
			"msg": "welcome PUT",
		})

		return
	case http.MethodDelete:
		c.JSON(http.StatusOK, gin.H{
			"msg": "welcome DELETE",
		})

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"panic": false,
		"name":  "METHOD_NOT_ALLOWED",
		"msg":   c.Request.Method + " method not allowed in this endpoint",
	})
}

func x(c *gin.Context) {
	var payload Login
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"panic": true,
			"name":  "REQUIRED_FIELDS",
			"msg":   "user and pass field is required",
		})

		return
	}

	if payload.User == "test" && payload.Pass == "123" {
		c.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"panic": false,
			"name":  "LOGIN_FAILED",
			"msg":   "please check your user and pass",
		})
	}
}
