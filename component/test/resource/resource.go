package resource

import (
	"github.com/javinc/mango/component/test/schema"
	"github.com/javinc/mango/database"

	db "github.com/gorethink/gorethink"
)

const resourceName = "test"

// Db instance
func Db() db.Term {
	return db.Table(resourceName)
}

// Find resource
func Find(o schema.Option) (data []schema.Object, err error) {
	q := baseFind(o)

	res, err := q.Run(database.Session)
	if err != nil {
		return
	}

	err = res.All(&data)
	if err != nil {
		return
	}

	// data being null if empty instead of array
	if data == nil {
		data = []schema.Object{}
	}

	return
}

func baseFind(o schema.Option) db.Term {
	return Db().Filter(map[string]interface{}{})
}
