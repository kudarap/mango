package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/javinc/puto/db"
	"github.com/javinc/puto/index"
	"github.com/javinc/puto/test"
)

// port server
const port = "8000"

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", index.Handler)
	r.HandleFunc("/test", test.Handler)

	// Migrates Db
	db.MySQL.AutoMigrate(&test.Model)

	// Bind to a port and pass our router in
	log.Println("serving on port", port)
	http.ListenAndServe(":"+port, r)
}
