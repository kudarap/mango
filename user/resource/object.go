package resource

import "time"

const objectName = "user"

// Model resource object
type Model struct {
	ID        int        `json:"id,omitempty"`
	First     string     `json:"first,omitempty" sql:"size:255"`
	Last      string     `json:"last,omitempty" sql:"size:255"`
	Email     string     `json:"email,omitempty" sql:"size:255"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_a,omitemptyt"`
	DeletedAt *time.Time `json:"-"`
}

// Options resource
type Options struct {
	Filters Model
	Fields  []string
	Page    struct {
		Limit  int
		Offset int
	}
	Sort struct {
		Asc  string
		Desc string
	}
}

// TableName specify table name required by ORM
func (x Model) TableName() string {
	return objectName
}

// LoadDefaults will load values if not modified
func (o *Options) LoadDefaults() {
	// fields defaults
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}

	// page limit & offset
	if o.Page.Limit == 0 {
		o.Page.Limit = -1
	}

	if o.Page.Offset == 0 {
		o.Page.Offset = -1
	}

	// sort
	if o.Sort.Desc != "" {
		o.Sort.Desc += " desc"
	}
}
