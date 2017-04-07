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
	return Errors{
		Name:    name,
		Message: message,
	}
}

// Panic returns generic custom error with panic
func Panic(name, message string) Errors {
	return Errors{
		Name:    name,
		Message: message,
		Panic:   true,
	}
}

// NewError accepts error value as a message
func NewError(name string, err error) Errors {
	return Errors{
		Name:    name,
		Message: err.Error(),
	}
}
