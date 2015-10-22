package test

import (
	"errors"
	r "github.com/javinc/puto/resources"
)

// Find resource
func Find(o Options) ([]Model, error) {
	models := []Model{}
	r.SQL.Find(&models, o.Filters)

	return models, nil
}

// Get resource
func Get(o Options) (Model, error) {
	model := Model{}
	r.SQL.First(&model, o.Filters)

	return model, nil
}

// Create resource
func Create(m *Model) (Model, error) {
	r.SQL.Create(&m)

	return *m, nil
}

// Remove resource
func Remove(i int) error {
	return errors.New("Remove method not available")
}

// Update resource
func Update(i int) error {
	return errors.New("Update method not available")
}
