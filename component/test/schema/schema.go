package schema

import "time"

// Object test
type Object struct {
	ID          string `json:"id,omitempty" field:"string"`
	UserID      string `json:"user_id,omitempty" field:"string"`
	Title       string `json:"title,omitempty" field:"string,required"`
	Description string `json:"description,omitempty" field:"string"`
	Age         int    `json:"age,omitempty" field:"number,indexed"`
	Taken       *bool  `json:"taken,omitempty" field:"boolean"`
	Hidden      *bool  `json:"hidden,omitempty" field:"boolean"`

	// meta
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Option test
type Option struct {
	Slice   string
	Order   string
	Search  string
	Field   string
	Filter  Object
	Between struct {
		Age string `index:"age"`
	}

	// RawFIlter is used for selective field search
	RawFilter map[string]interface{}
}
