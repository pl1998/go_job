/**
  @author:panliang
  @data:2021/8/4
  @note
**/
package db

import (
	"go_job/pkg/helpers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Conn()  {
	DB = like()
}

func like() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("/Users/panliang/www/admin_api/database/admin_api.db"), &gorm.Config{})
	helpers.CheckErr(err)
	return db
}
