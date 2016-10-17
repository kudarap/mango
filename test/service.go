package test

// Service test
type Service struct {
}

var resource Resource

// Find test
func (t *Service) Find() []Object {
	return resource.Find()
}

// Get test
func (t *Service) Get() {
}

// Create test
func (t *Service) Create() {
}

// Update test
func (t *Service) Update() {
}

// Remove test
func (t *Service) Remove() {
}
