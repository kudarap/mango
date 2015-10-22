package service

import r "github.com/javinc/puto/test/resource"

// Create service
func Create(m *r.Model) r.Model {
	row, _ := r.Create(m)

	return row
}
