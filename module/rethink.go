package module

import (
	"errors"
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

// Rethink inialiization
func Rethink() {
	var err error

	// check credentials
	if Config.Rethink.Host == "" || Config.Rethink.Db == "" {
		log.Fatalln(errors.New("rethink host or db not defined"))
	}

	RSession, err = r.Connect(r.ConnectOpts{
		Address:  Config.Rethink.Host,
		Database: Config.Rethink.Db,
		MaxOpen:  Config.Rethink.MaxOpen,
	})

	if err != nil {
		log.Fatalln(err)

		return
	}

	r.DBCreate(Config.Rethink.Db).Run(RSession)
}
