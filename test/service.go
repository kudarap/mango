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
func (t *Service) Get(id string) Object {
	return resource.Get(id)
}

// Create test
func (t *Service) Create(p Object) Object {
	return resource.Create(p)
}

// Update test
func (t *Service) Update(p Object, id string) Object {
	return resource.Update(p, id)
}

// Remove test
func (t *Service) Remove(id string) bool {
	return resource.HardRemove(id)
}
