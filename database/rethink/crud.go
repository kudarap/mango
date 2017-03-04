package rethink

import "github.com/gorethink/gorethink"

// Run query
func Run(term gorethink.Term) (*gorethink.Cursor, error) {
	return term.Run(GetSession())
}

// Find query
func Find(term gorethink.Term, result interface{}) error {
	r, err := Run(term)
	if err != nil {
		return err
	}

	err = r.All(result)
	if err != nil {
		return err
	}

	return nil
}
