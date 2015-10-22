package service

import r "github.com/javinc/puto/test/resource"

// Find service
func Find(o *r.Options) []r.Model {
	rows, _ := r.Find(*o)

	return rows
}
