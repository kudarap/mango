package test

import (
	"log"
	"time"

	r "github.com/dancannon/gorethink"
	"github.com/javinc/mango/module"
)

const tableName = "test"

// Resource test
type Resource struct {
}

// Object test
type Object struct {
	ID          string     `gorethink:"id,omitempty" json:"id"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
}

// Find test
func (t *Resource) Find() []Object {
	data := []Object{}

	res, err := r.Table(tableName).Run(module.RSession)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = res.All(&data)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return data
}

// Get test
func (t *Resource) Get() Object {
	return Object{
		ID:          "100",
		Title:       "James",
		Description: "testing ground",
	}
}

// Create test
func (t *Resource) Create() {
	// Create the item
	payload := Object{
		Title:       "James",
		Description: "testing ground",
	}

	// meta data
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	// insert to database
	_, err := r.Table(tableName).Insert(payload).RunWrite(module.RSession)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}

// Update test
func (t *Resource) Update() {
}

// Remove test
func (t *Resource) Remove() {
}
