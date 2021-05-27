package common

import (
	"database/sql"
	"hkzhao/go_blog_system/config"
)

func InitMySQL() (d *sql.DB, err error) {
	var db *sql.DB
	dsn := config.MYSQLURL
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
