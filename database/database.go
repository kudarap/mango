package database

import (
	"errors"
	"log"

	db "github.com/gorethink/gorethink"

	"github.com/javinc/mango/config"
)

// Config database
type Config struct {
	Host    string
	Db      string
	MaxOpen int
}

var s *db.Session

func init() {
	log.Println("[database]", "starting...")

	c := config.Get()
	Connect(Config{
		Host:    c.Rethink.Host,
		Db:      c.Rethink.Db,
		MaxOpen: c.Rethink.MaxOpen,
	})

	log.Println("[database]", "started!")
}

// GetSession database
func GetSession() *db.Session {
	return s
}

// Connect to database
func Connect(c Config) {
	// check credentials
	err := errors.New("db host or db must defined")
	if c.Host == "" || c.Db == "" {
		log.Fatalln(err)
	}

	s, err = db.Connect(db.ConnectOpts{
		Address:  c.Host,
		Database: c.Db,
		MaxOpen:  c.MaxOpen,
	})
	if err != nil {
		log.Fatalln("[database]", err)
	}

	// create if not exists
	db.DBCreate(c.Db).Run(s)

	// enabling json tag as alternative on component Objects
	db.SetTags("gorethink", "json")
}
