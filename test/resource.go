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
func (t *Resource) Get(id string) Object {
	data := Object{}

	res, err := r.Table(tableName).Get(id).Run(module.RSession)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = res.One(&data)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return data
}

// Create test
func (t *Resource) Create(p Object) Object {
	// Create the item
	payload := p

	// meta data
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	// insert to database
	_, err := r.Table(tableName).Insert(payload).RunWrite(module.RSession)
	if err != nil {
		log.Fatalln(err.Error())

		return Object{}
	}

	return payload
}

// Update test
func (t *Resource) Update() {

}

// Remove test
func (t *Resource) Remove() {
}
