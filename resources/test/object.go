package test

import "time"

const objectName = "test"

// object model
type Model struct {
    Id int `json:"id"`
    Title string `json:"title" sql:"size:255"`
    Description string `json:"description" sql:"size:255"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"-"`
}

type Options struct {
    Filters Model
    Fields []string
    Limits []int
    Sorts struct {
        Asc []string
        Desc []string
    }
}

// specify table name required by ORM
func (x Model) TableName() string {
    return objectName
}
