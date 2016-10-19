package module

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var contcext *gin.Context

// SetContext format
func SetContext(c *gin.Context) {
	contcext = c
}

// Error format
func Error(name string, msg string) {
	contcext.JSON(http.StatusBadRequest, gin.H{
		"panic": false,
		"name":  name,
		"msg":   msg,
	})
}

// Panic format
func Panic(name string, msg string) {
	contcext.JSON(http.StatusInternalServerError, gin.H{
		"panic": true,
		"name":  name,
		"msg":   msg,
	})
}

// Output format
func Output(data interface{}) {
	contcext.JSON(http.StatusOK, data)
}
