package module

import (
	"log"

	r "github.com/dancannon/gorethink"
)

// RSession rethink
var (
	RSession *r.Session
)

func init() {
	var err error

	RSession, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "mango",
		MaxOpen:  40,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}
}
