package model

import "time"

// User test model
type User struct {
	ID        string     `json:"id,omitempty" field:"string"`
	Type      string     `json:"type,omitempty" field:"string"`
	Email     string     `json:"email,omitempty" field:"string"`
	Active    *bool      `json:"taken,omitempty" field:"boolean"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Foo test model
type Foo struct {
	ID          string     `json:"id,omitempty" field:"string"`
	UserID      string     `json:"user_id,omitempty" field:"string"`
	Title       string     `json:"title,omitempty" field:"string,required"`
	Description string     `json:"description,omitempty" field:"string"`
	Age         int        `json:"age,omitempty" field:"number,indexed"`
	Taken       *bool      `json:"taken,omitempty" field:"boolean"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
