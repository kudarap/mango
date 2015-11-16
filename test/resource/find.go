package resource

import x "github.com/javinc/puto"

// Find resource
func Find(o Options) ([]Model, error) {
	models := []Model{}
	e := x.MySQL.Find(&models, o.Filters).Error

	return models, e
}

// Get resource
func Get(o Options) (Model, error) {
	model := Model{}
	e := x.MySQL.First(&model, o.Filters).Error

	return model, e
}
