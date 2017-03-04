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

// Init connection
func Init() {
	log.Println("[database]", "starting...")

	log.Println("[database]", "started!")
}

// GetSession database
func GetSession() *gorethink.Session {
	return s
}

// Connect to database
func Connect(c Config) {
	// check credentials
	err := errors.New("db host or db must defined")
	if c.Host == "" || c.Db == "" {
		log.Fatalln(err)
	}

	s, err = gorethink.Connect(gorethink.ConnectOpts{
		Address:  c.Host,
		Database: c.Db,
		MaxOpen:  c.MaxOpen,
	})
	if err != nil {
		log.Fatalln("[database]", err)
	}

	// create if not exists
	gorethink.DBCreate(c.Db).Run(s)

	// enabling json tag as alternative on component Objects
	gorethink.SetTags("gorethink", "json")
}
