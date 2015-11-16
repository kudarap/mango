package resource

import x "github.com/javinc/puto"

// Create resource
func Create(m *Model) (Model, error) {
	e := x.MySQL.Create(&m).Error

	return *m, e
}
