package module

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var context *gin.Context

// SetContext format
func SetContext(c *gin.Context) {
	context = c
}

// Error format
func Error(name string, msg string) {
	context.JSON(http.StatusBadRequest, gin.H{
		"panic": false,
		"name":  name,
		"msg":   msg,
	})
}

// Panic format
func Panic(name string, msg string) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"panic": true,
		"name":  name,
		"msg":   msg,
	})
}

// Output format
func Output(data interface{}) {
	context.JSON(http.StatusOK, data)
}
