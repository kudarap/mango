package service

import r "github.com/javinc/puto/test/resource"

// Update service
func Update(m *r.Model, o *r.Options) (r.Model, error) {
	row, err := r.Update(*m, *o)
	if err != nil {
		return r.Model{}, err
	}

	return row, err
}
