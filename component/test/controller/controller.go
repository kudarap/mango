package controller

import (
	"github.com/javinc/mango/component/test/schema"
	"github.com/javinc/mango/component/test/service"
)

// Handler http
func Handler() {
	o := schema.Option{}

	service.Find(o)
}
