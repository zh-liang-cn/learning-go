package common

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DEBUG = true

var _db *gorm.DB

func InitDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("tiny-wiki.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	_db = db
	return
}

func GetDB() *gorm.DB {
	if DEBUG {
		return _db.Debug()
	} else {
		return _db
	}
}
