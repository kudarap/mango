package resource

import x "github.com/javinc/puto"

// Create resource
func Create(m *Model) (Model, error) {
	x.MySQL.Create(&m)

	return *m, nil
}
