package resource

import (
	"errors"

	x "github.com/javinc/puto"
)

// Search resource
func Search(key string) ([]Model, error) {
	models := []Model{}

	e := x.MySQL.
		Where("title LIKE ?", "%"+key+"%").
		Or("description LIKE ?", "%"+key+"%").
		Find(&models).
		Error

	return models, e
}

// Find resource
func Find(o Options) ([]Model, error) {
	models := []Model{}

	e := x.MySQL.
		Select(o.Fields).
		Limit(o.Page.Limit).
		Offset(o.Page.Offset).
		Order(o.Sort.Asc).
		Order(o.Sort.Desc).
		Find(&models, o.Filters).
		Error

	return models, e
}

// Get resource
func Get(o Options) (Model, error) {
	// set limit to 1 and offset to default
	// to prevent overrides
	o.Page.Limit = 1
	o.Page.Offset = -1

	models, e := Find(o)

	// check if exists
	if len(models) == 0 {
		return Model{}, errors.New("record not found")
	}

	// get single
	return models[0], e
}
