package test

// Service test
type Service struct {
}

var resource Resource

// Find test
func (t *Service) Find(o Option) ([]Object, error) {
	return resource.Find(o)
}

// Get test
func (t *Service) Get(id string) (Object, error) {
	return resource.Get(id)
}

// Create test
func (t *Service) Create(p Object) (Object, error) {
	return resource.Create(p)
}

// Update test
func (t *Service) Update(p Object, id string) (Object, error) {
	return resource.Update(p, id)
}

// Remove test
func (t *Service) Remove(id string) (bool, error) {
	return resource.Remove(id)
}
