package resource

import x "github.com/javinc/puto"

// Find resource
func Find(o Options) ([]Model, error) {
	models := []Model{}

	// fields defaults
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}

	// limit & offset
	if len(o.Limits) == 0 {
		o.Limits = []int{-1, -1}
	} else if len(o.Limits) == 1 {
		o.Limits = []int{o.Limits[0], -1}
	}

	// sort
	if o.Sorts.Desc != "" {
		o.Sorts.Desc += " desc"
	}

	e := x.MySQL.
		Select(o.Fields).
		Limit(o.Limits[0]).
		Offset(o.Limits[1]).
		Order(o.Sorts.Asc).
		Order(o.Sorts.Desc).
		Find(&models, o.Filters).
		Error

	return models, e
}

// Get resource
func Get(o Options) (Model, error) {
	model := Model{}
	e := x.MySQL.First(&model, o.Filters).Error

	return model, e
}
