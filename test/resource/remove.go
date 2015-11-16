package resource

import x "github.com/javinc/puto"

// Remove resource
func Remove(o Options) (Model, error) {
	model, _ := Get(o)
	x.MySQL.Delete(&model)

	return model, nil
}
