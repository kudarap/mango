package errors

import "fmt"

// Errors model
type Errors struct {
	Name    string      `json:"name,omitempty"`
	Message string      `json:"message,omitempty"`
	Panic   bool        `json:"panic,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (e Errors) Error() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}

// New returns generic custom error
func New(name, message string) Errors {
	return create(name, message, false)
}

// NewError accepts error value as a message
func NewError(name string, err error) Errors {
	return create(name, err.Error(), false)
}

// Panic returns generic custom error with panic
func Panic(name, message string) Errors {
	return create(name, message, true)
}

// PanicError returns generic custom error with panic and accepts error value
func PanicError(name string, err error) Errors {
	return create(name, err.Error(), true)
}

func create(n, m string, p bool) Errors {
	return Errors{
		Name:    n,
		Message: m,
		Panic:   p,
	}
}
