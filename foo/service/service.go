package service

import (
	"github.com/javinc/mango/foo/resource"
	"github.com/javinc/mango/foo/schema"
)

// Find test
func Find(o schema.Option) ([]schema.Object, error) {
	return resource.Find(o)
}
