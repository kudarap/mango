package test

import (
    "errors"
    r "github.com/javinc/playgo/goryo/resources"
)

func Find(o Options) ([]Model, error) {
    models := []Model{}
    r.Sql.Find(&models, o.Filters)

    return models, nil
}

func Get(o Options) (Model, error) {
    model := Model{}
    r.Sql.First(&model, o.Filters)

    return model, nil
}

func Create(m *Model) (Model, error) {
    r.Sql.Create(&m)

    return *m, nil
}

func Remove(i int) error {
    return errors.New("Remove method not available")
}

func Update(i int) error {
    return errors.New("Update method not available")
}
