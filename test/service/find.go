package service

import r "github.com/javinc/puto/test/resource"

// Find service
func Find(o *r.Options) ([]r.Model, error) {
	rows, err := r.Find(*o)
	if err != nil {
		return []r.Model{}, err
	}

	return rows, err
}

// Get service
func Get(o *r.Options) (r.Model, error) {
	row, err := r.Get(*o)
	if err != nil {
		return r.Model{}, err
	}

	return row, err
}
