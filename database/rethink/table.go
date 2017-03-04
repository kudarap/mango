package rethink

import "github.com/gorethink/gorethink"

// CreateTable create table
func CreateTable(name string) {
	gorethink.TableCreate("foo").
		Wait().
		RunWrite(GetSession())
}
