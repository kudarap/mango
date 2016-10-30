package module

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ctx *gin.Context

// Router gin
func Router() *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}

// SetContext format
func SetContext(c *gin.Context) {
	ctx = c
}

// Error format
func Error(name string, msg string) {
	checkContext()

	ctx.JSON(http.StatusBadRequest, gin.H{
		"panic": false,
		"name":  name,
		"msg":   msg,
	})
}

// MethodNotAllowed restricted method
func MethodNotAllowed() {
	checkContext()

	Error("METHOD_NOT_ALLOWED",
		ctx.Request.Method+" method not allowed in this endpoint")
}

// Panic format
func Panic(name string, msg string) {
	checkContext()

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"panic": true,
		"name":  name,
		"msg":   msg,
	})
}

// Output format
func Output(data interface{}) {
	checkContext()

	ctx.JSON(http.StatusOK, data)
}

// GET format
func GET(handler func()) {
	checkContext()

	if ctx.Request.Method == http.MethodGet {
		handler()
	}
}

func checkContext() {
	if ctx == nil {
		log.Fatalln("ctx is undefined, please set it on handler")
	}
}
