package rethink

import (
	"errors"
	"log"

	"github.com/gorethink/gorethink"
)

// Config database
type Config struct {
	Host    string
	Db      string
	MaxOpen int
}

var s *gorethink.Session

// GetSession database
func GetSession() *gorethink.Session {
	return s
}

// Connect to database
func Connect(c Config) {
	// check credentials
	err := errors.New("[database] Config host, db, and maxOpen must defined")
	if c.Host == "" || c.Db == "" || c.MaxOpen == 0 {
		log.Println(err)
	}

	s, err = gorethink.Connect(gorethink.ConnectOpts{
		Address:  c.Host,
		Database: c.Db,
		MaxOpen:  c.MaxOpen,
	})
	if err != nil {
		log.Println("[database]", err)
	}

	// create if not exists
	gorethink.DBCreate(c.Db).Run(s)

	// enabling json tag as alternative on component Objects
	gorethink.SetTags("gorethink", "data")
}
