package rethink

import "github.com/gorethink/gorethink"

// Run query
func Run(term gorethink.Term) (*gorethink.Cursor, error) {
	return term.Run(GetSession())
}

// RunWrite query
func RunWrite(term gorethink.Term) (gorethink.WriteResponse, error) {
	return term.RunWrite(GetSession())
}
