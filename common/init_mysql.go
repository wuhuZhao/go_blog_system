package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hkzhao/go_blog_system/config"
)

var Db *gorm.DB

func InitMySQL() (err error) {
	dsn := config.MYSQLURL
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Db = d
	return nil
}
