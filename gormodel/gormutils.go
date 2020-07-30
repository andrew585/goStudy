package gormodel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var orm *gorm.DB

func Init() error {
	db, err := gorm.Open("mysql",
		"root:root@/studydb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	orm = db
	return nil
}
