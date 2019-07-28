package db

import (
	"dashboard/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.db
	err error
)

func Init() {
	db, err = gorm.Open("mysql", "host=localhost port=5432 user=root dbname=dashboard")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
