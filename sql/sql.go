package sql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Item struct {
	gorm.Model
	Name        string    `form:"name"`
	Type        string    `form:"type"`
	Count       int       `form:"count"`
	ProduceDate time.Time `form:"produce_date" time_format:"2006-01-02" time_utc:"8"`
	SafeDay     int       `form:"safe_day"`
}

func Test() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Item{})
	if err != nil {
		return
	}
	return db
}
