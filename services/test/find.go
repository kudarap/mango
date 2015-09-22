package test

import "github.com/javinc/playgo/goryo/resources/test"

func Find(o *test.Options) []test.Model {
    rows, _ := test.Find(*o)

    return rows
}
