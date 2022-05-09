package sql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"timedo/common"
)

func Test() (db *gorm.DB) {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&common.Item{})
	if err != nil {
		return
	}
	return db
}
