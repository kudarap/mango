package service

import (
	"github.com/javinc/mango/component/test/resource"
	"github.com/javinc/mango/component/test/schema"
)

// Find test
func Find(o schema.Option) ([]schema.Object, error) {
	return resource.Find(o)
}
