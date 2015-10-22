package test

import "github.com/javinc/puto/resources/test"

// Find service
func Find(o *test.Options) []test.Model {
	rows, _ := test.Find(*o)

	return rows
}
