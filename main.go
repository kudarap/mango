package main

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/javinc/playgo/goryo/resources"
    "github.com/javinc/playgo/goryo/resources/test"
    // "github.com/javinc/playgo/goryo/resources/user"
    "github.com/javinc/playgo/goryo/controllers"
)

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", controllers.IndexHandler)
    r.HandleFunc("/test", controllers.TestHandler)
    r.HandleFunc("/user", controllers.UserHandler)

    // Migrates Db
    resources.Sql.AutoMigrate(&test.Model{})
    // resources.Sql.AutoMigrate(&user.Model{})

    // Bind to a port and pass our router in
    http.ListenAndServe(":8000", r)
}
