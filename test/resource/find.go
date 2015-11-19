package resource

import x "github.com/javinc/puto"

// Find resource
func Find(o Options) ([]Model, error) {
	models := []Model{}

	// check fields if blank make it *
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}

	e := x.MySQL.Select(o.Fields).Find(&models, o.Filters).Error

	return models, e
}

// Get resource
func Get(o Options) (Model, error) {
	model := Model{}
	e := x.MySQL.First(&model, o.Filters).Error

	return model, e
}
