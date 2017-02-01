package controller

import (
	"github.com/javinc/mango/foo/schema"
	"github.com/javinc/mango/foo/service"
)

// Handler http
func Handler() {
	o := schema.Option{}

	service.Find(o)
}
