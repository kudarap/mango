package resource

import r "github.com/javinc/puto/db"

// Find resource
func Find(o Options) ([]Model, error) {
	models := []Model{}
	r.MySQL.Find(&models, o.Filters)

	return models, nil
}

// Get resource
func Get(o Options) (Model, error) {
	model := Model{}
	r.MySQL.First(&model, o.Filters)

	return model, nil
}
