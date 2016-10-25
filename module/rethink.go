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

	host := "localhost"
	port := "28015"
	db := "mango"
	maxOpen := 40

	RSession, err = r.Connect(r.ConnectOpts{
		Address:  host + ":" + port,
		Database: db,
		MaxOpen:  maxOpen,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	r.DBCreate(db).Run(RSession)
}
