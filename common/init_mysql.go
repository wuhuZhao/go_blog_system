package common

import (
	"database/sql"
	"hkzhao/go_blog_system/config"
)

var db *sql.DB

func initMySQL() (err error) {
	dsn := config.MYSQLURL
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
