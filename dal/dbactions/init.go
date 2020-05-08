package dbactions

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@/mememe?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.DB()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(false)

	db.LogMode(true)
}
