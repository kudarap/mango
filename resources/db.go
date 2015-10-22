package resources

import (
	"log"

	// mysql driver imported
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQL database instance
var SQL = Db()

const (
	db   = "puto"
	user = "root"
	pass = "root"
)

// Db instantiation
func Db() gorm.DB {
	x, err := gorm.Open("mysql",
		user+":"+pass+"@/"+db+"?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	x.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	x.DB().Ping()
	x.DB().SetMaxIdleConns(1000)
	x.DB().SetMaxOpenConns(1000)

	// Disable table name's pluralization
	x.SingularTable(true)
	x.LogMode(true)

	return x
}
