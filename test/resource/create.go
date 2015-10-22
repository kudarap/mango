package resource

import r "github.com/javinc/puto/db"

// Create resource
func Create(m *Model) (Model, error) {
	r.MySQL.Create(&m)

	return *m, nil
}
