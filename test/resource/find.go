package resource

import x "github.com/javinc/puto"

// Find resource
func Find(o Options) ([]Model, error) {
	models := []Model{}
	x.MySQL.Find(&models, o.Filters)

	return models, nil
}

// Get resource
func Get(o Options) (Model, error) {
	model := Model{}
	x.MySQL.First(&model, o.Filters)

	return model, nil
}
