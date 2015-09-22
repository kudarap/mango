package resources

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var Sql = Db()

const (
	db = "ejr"
	user = "root"
	pass = "root"
)

func Db() gorm.DB {
    x, err := gorm.Open("mysql",
        user + ":" + pass + "@/" + db + "?charset=utf8&parseTime=True")
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
