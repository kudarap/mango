package service

import r "github.com/javinc/puto/test/resource"

// Update service
func Update(m *r.Model, o *r.Options) r.Model {
	row, _ := r.Update(*m, *o)

	return row
}
