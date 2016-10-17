package module

import (
	"log"

	r "github.com/dancannon/gorethink"
)

// RSession rethink
var (
	RSession *r.Session

	host     = "localhost"
	port     = "28015"
	database = "mango"
	maxOpen  = 40
)

func init() {
	var err error

	RSession, err = r.Connect(r.ConnectOpts{
		Address:  host + ":" + port,
		Database: database,
		MaxOpen:  maxOpen,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}
}
