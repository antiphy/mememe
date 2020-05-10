package dbactions

import (
	"github.com/antiphy/mememe/dal/consts"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", consts.GetDBSourceName())
	if err != nil {
		panic(err)
	}

	db.DB()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(false)
	db.LogMode(true)
}
