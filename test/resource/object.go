package resource

import "time"

const objectName = "test"

// Model resource object
type Model struct {
	ID          int        `json:"id"`
	Title       string     `json:"title" sql:"size:255"`
	Description string     `json:"description" sql:"size:255"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
}

// Options resource
type Options struct {
	Search  string
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
