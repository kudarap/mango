package resource

import x "github.com/javinc/puto"

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

	// get 1st
	return models[0], e
}
