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
	Filters Model
	Fields  []string
	Limits  []int
	Sorts   struct {
		Asc  []string
		Desc []string
	}
}

// TableName specify table name required by ORM
func (x Model) TableName() string {
	return objectName
}

// LoadDefaults will load values if not modified
func (o Options) LoadDefaults() {
	// fields if blank make it *
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}

}
