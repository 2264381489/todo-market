package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitModel(datasource string) *gorm.DB {
	//datasource = "root:123456@(127.0.0.1:13306)/dev?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	return db
}
