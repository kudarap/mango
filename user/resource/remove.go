package resource

import x "github.com/javinc/puto"

// Remove resource
func Remove(o Options) (Model, error) {
	model, e := Get(o)
	if e != nil {
		return Model{}, e
	}

	e = x.MySQL.Delete(&model).Error

	return model, e
}
