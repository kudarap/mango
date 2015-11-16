package service

import r "github.com/javinc/puto/test/resource"

// Remove service
func Remove(o *r.Options) r.Model {
	row, err := r.Remove(*o)

	return row
}
