package rethink

import "github.com/gorethink/gorethink"

// CreateTable create table
func CreateTable(name string) {
	gorethink.TableCreate("foo").
		Wait().
		RunWrite(GetSession())
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

// FindOne query
func FindOne(term gorethink.Term, result interface{}) error {
	r, err := Run(term)
	if err != nil {
		return err
	}

	err = r.One(result)
	if err != nil {
		return err
	}

	return nil
}

// Create query
func Create(term gorethink.Term) (string, error) {
	r, err := RunWrite(term)
	if err != nil {
		return "", err
	}

	// it should return atleast 1 key
	return r.GeneratedKeys[0], nil
}
