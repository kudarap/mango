package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/javinc/puto/controllers"
	"github.com/javinc/puto/resources"
	"github.com/javinc/puto/resources/test"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", controllers.IndexHandler)
	r.HandleFunc("/test", controllers.TestHandler)

	// Migrates Db
	resources.SQL.AutoMigrate(&test.Model{})

	// Bind to a port and pass our router in
	http.ListenAndServe(":8000", r)
}
