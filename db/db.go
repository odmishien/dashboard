package db

import (
	"dashboard/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/dashboard")
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
	db.AutoMigrate(&entity.Song{})
}
