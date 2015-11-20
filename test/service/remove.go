package service

import r "github.com/javinc/puto/test/resource"

// Remove service
func Remove(o *r.Options) (r.Model, error) {
	row, err := r.Remove(*o)
	if err != nil {
		return r.Model{}, err
	}

	return row, err
}
