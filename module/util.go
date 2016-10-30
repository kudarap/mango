package module

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type config struct {
	Host    string `json:"host"`
	Rethink rethinkConfig
}

const (
	// GET method
	GET = http.MethodGet
	// POST method
	POST = http.MethodPost
	// PATCH method
	PATCH = http.MethodPatch
	// DELETE method
	DELETE = http.MethodDelete

	configPath = "config.json"
)

var (
	ctx *gin.Context
	// Config settings
	Config config
)

func init() {
	file, _ := os.Open(configPath)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Config)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("config.json loaded")
	log.Println(Config)
}

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

func checkContext() {
	if ctx == nil {
		log.Fatalln("ctx is undefined, please set it on handler")
	}
}
