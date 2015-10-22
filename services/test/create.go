package test

import "github.com/javinc/puto/resources/test"

// Create service
func Create(m *test.Model) test.Model {
	row, _ := test.Create(m)

	return row
}
