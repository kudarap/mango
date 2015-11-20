package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	x "github.com/javinc/puto"
	"github.com/javinc/puto/index"
	"github.com/javinc/puto/test"
	"github.com/javinc/puto/user"
)

// port server
const port = "8000"

func main() {
	m := mux.NewRouter()

	// Routes consist of a path and a handler function.
	m.HandleFunc("/", index.Handler)
	m.HandleFunc("/test", test.Handler)
	m.HandleFunc("/test/{id:[0-9]+}", test.Handler)
	m.HandleFunc("/user", user.Handler)
	m.HandleFunc("/user/{id:[0-9]+}", user.Handler)

	// Migrates Db
	x.MySQL.AutoMigrate(
		&test.Model,
		&user.Model,
	)

	// Bind to a port and pass our router in
	log.Println("serving on port", port)
	http.ListenAndServe(":"+port, m)
}
