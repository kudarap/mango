package rethink

import "github.com/gorethink/gorethink"

// Find basic query
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

// FindOne basic query
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

// Create basic query
func Create(term gorethink.Term) (string, error) {
	r, err := RunWrite(term)
	if err != nil {
		return "", err
	}

	// it should return atleast 1 key
	return r.GeneratedKeys[0], nil
}

// Update basic query
func Update(term gorethink.Term) error {
	_, err := RunWrite(term)
	if err != nil {
		return err
	}

	return nil
}

// CreateTable create table
func CreateTable(name string) {
	gorethink.TableCreate("foo").
		Wait().
		RunWrite(GetSession())
}
