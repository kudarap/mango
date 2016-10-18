package test

import (
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
}

// Find test
func (t *Resource) Find() ([]Object, error) {
	data := []Object{}

	res, err := r.Table(tableName).
		OrderBy(r.Desc("created_at")).
		Run(module.RSession)
	if err != nil {
		return data, err
	}

	err = res.All(&data)
	if err != nil {
		return data, err
	}

	return data, err
}

// Get test
func (t *Resource) Get(id string) (Object, error) {
	data := Object{}

	res, err := r.Table(tableName).Get(id).Run(module.RSession)
	if err != nil {
		return data, err
	}

	err = res.One(&data)
	if err != nil {
		return data, err
	}

	return data, err
}

// Create test
func (t *Resource) Create(p Object) (Object, error) {
	// meta data
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	// insert to database
	_, err := r.Table(tableName).Insert(p).Run(module.RSession)
	if err != nil {
		return Object{}, err
	}

	return p, err
}

// Update test
func (t *Resource) Update(p Object, id string) (Object, error) {
	// check item if exists
	_, err := t.Get(id)
	if err != nil {
		return Object{}, err
	}

	// update meta data
	p.UpdatedAt = time.Now()

	// insert to database
	_, err = r.Table(tableName).Get(id).Update(p).Run(module.RSession)
	if err != nil {
		return Object{}, err
	}

	return p, err
}

// Remove test
func (t *Resource) Remove(id string) (bool, error) {
	// check item if exists
	_, err := t.Get(id)
	if err != nil {
		return false, err
	}

	// insert to database
	_, err = r.Table(tableName).Get(id).Delete().Run(module.RSession)
	if err != nil {
		return false, err
	}

	return true, nil
}
