package resource

import (
	"log"

	"github.com/gorethink/gorethink"

	"github.com/javinc/mango/database"
	"github.com/javinc/mango/foo/schema"
)

const resourceName = "foo"

var session = database.GetSession()

func init() {
	log.Println("[foo]", "init")

	gorethink.TableCreate(resourceName).Wait().RunWrite(session)
}

// Db instance
func Db() gorethink.Term {
	return gorethink.Table(resourceName)
}

// Find resource
func Find(o schema.Option) (data []schema.Object, err error) {
	q := baseFind(o)

	res, err := q.Run(session)
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

func baseFind(o schema.Option) gorethink.Term {
	return Db().Filter(map[string]interface{}{})
}
