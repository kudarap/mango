package module

import (
	"log"

	r "github.com/dancannon/gorethink"
)

type rethinkConfig struct {
	Host    string `json:"host"`
	Db      string `json:"db"`
	MaxOpen int    `json:"max_open"`
}

// RSession rethink
var (
	RSession *r.Session
)

func init() {
	var err error

	RSession, err = r.Connect(r.ConnectOpts{
		Address:  Config.Rethink.Host,
		Database: Config.Rethink.Db,
		MaxOpen:  Config.Rethink.MaxOpen,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	r.DBCreate(Config.Rethink.Db).Run(RSession)
}
