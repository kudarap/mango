package puto

import (
	"log"

	// mysql driver imported
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// MySQL database instance
var MySQL = Db()

const (
	db   = "puto"
	user = "root"
	pass = "root"
)

// Db instantiation
func Db() gorm.DB {
	d, err := gorm.Open("mysql",
		user+":"+pass+"@/"+db+"?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	// Get database connection handle
	// [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	d.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	d.DB().Ping()
	d.DB().SetMaxIdleConns(1000)
	d.DB().SetMaxOpenConns(1000)

	// Disable table name's pluralization
	d.SingularTable(true)
	d.LogMode(true)

	return d
}
