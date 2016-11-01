package module

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

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
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Before", "Foo")
		c.Next()
	})

	return r
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Before", "Foo")
		c.Next()
		c.Writer.Header().Set("X-After", "Bar")
	}
}

// SetContext format
func SetContext(c *gin.Context) {
	ctx = c
}

// MethodNotAllowedError restricted method
func MethodNotAllowedError() {
	checkContext()

	Error("METHOD_NOT_ALLOWED",
		ctx.Request.Method+" method not allowed in this endpoint")
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

// Panic format
func Panic(name string, msg string) {
	checkContext()

	ctx.Header("content-type", "application/json")
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

// Hash using MD5
func Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text + appKey))

	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateHash using Hash
func GenerateHash() string {
	return Hash(time.Now().String())
}

func checkContext() {
	if ctx == nil {
		log.Fatalln("ctx is undefined, please set it on handler")
	}
}
