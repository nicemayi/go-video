package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(157.230.169.141:3306)/videoserver?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
