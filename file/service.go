package file

// Service file
type Service struct {
}

var resource Resource

// Find file
func (t *Service) Find(o Option) ([]Object, error) {
	return resource.Find(o)
}

// Get file
func (t *Service) Get(id string) (Object, error) {
	return resource.Get(id)
}

// Create file
func (t *Service) Create(p Object) (Object, error) {
	return resource.Create(p)
}

// Update file
func (t *Service) Update(p Object, id string) (Object, error) {
	return resource.Update(p, id)
}

// Remove file
func (t *Service) Remove(id string) (Object, error) {
	return resource.Remove(id)
}
