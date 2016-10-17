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
	ID          string    `gorethink:"id,omitempty" json:"id,omitempty"`
	Title       string    `gorethink:"title,omitempty" json:"title,omitempty"`
	Description string    `gorethink:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time `gorethink:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorethink:"updated_at,omitempty" json:"updated_at,omitempty"`
	Deleted     bool      `gorethink:"deleted,omitempty" json:"deleted"`
}

// Find test
func (t *Resource) Find() []Object {
	data := []Object{}

	res, err := r.Table(tableName).
		OrderBy(r.Desc("created_at")).
		Run(module.RSession)
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

	res, err := r.Table(tableName).
		Get(id).Run(module.RSession)
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
	// meta data
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Deleted = false

	// insert to database
	_, err := r.Table(tableName).Insert(p).Run(module.RSession)
	if err != nil {
		log.Fatalln(err.Error())

		return Object{}
	}

	return p
}

// Update test
func (t *Resource) Update(p Object, id string) Object {
	// check item if exists
	payload := t.Get(id)
	if payload.ID == "" {
		log.Fatalln("not exists")

		return Object{}
	}

	// meta data
	p.UpdatedAt = time.Now()

	// insert to database
	_, err := r.Table(tableName).Get(id).Update(p).Run(module.RSession)
	if err != nil {
		log.Fatalln(err.Error())

		return Object{}
	}

	return p
}

// Remove test
func (t *Resource) Remove(id string) bool {
	// check item if exists
	payload := t.Get(id)
	if payload.ID == "" {
		log.Fatalln("not exists")

		return false
	}

	// meta data
	p := Object{}
	p.UpdatedAt = time.Now()
	p.Deleted = true

	// insert to database
	_, err := r.Table(tableName).Get(id).Update(p).Run(module.RSession)
	if err != nil {
		log.Fatalln(err.Error())

		return false
	}

	return true
}

// HardRemove test
func (t *Resource) HardRemove(id string) bool {
	// check item if exists
	payload := t.Get(id)
	if payload.ID == "" {
		log.Fatalln("not exists")

		return false
	}

	// insert to database
	_, err := r.Table(tableName).Get(id).Delete().Run(module.RSession)
	if err != nil {
		log.Fatalln(err.Error())

		return false
	}

	return true
}
