package resource

import "log"

func init() {
	log.Println("[foo-store]", "init")
}

// GetFoo get foo
func (db *resource) GetFoo(id string) (string, error) {
	return "resource foo", nil
}
