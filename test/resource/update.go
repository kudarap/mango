package resource

import x "github.com/javinc/puto"

// Update resource
func Update(p Model, o Options) (Model, error) {
	model, e := Get(o)
	if e != nil {
		return Model{}, e
	}

	e = x.MySQL.Model(&model).Updates(p).Error

	return model, e
}
