package logic

import "log"

func init() {
	log.Println("[foo-service]", "init")
}

// GetFoo test
func (db *logic) GetFoo(id string) (string, error) {
	return "service foo", nil
}
