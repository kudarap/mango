package test

import "github.com/javinc/playgo/goryo/resources/test"

func Create(m *test.Model) test.Model {
    row, _ := test.Create(m)

    return row
}
