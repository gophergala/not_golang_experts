package conf

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func SetupDB() *gorm.DB {
	db, err := gorm.Open("postgres", "dbname=gopherstalker sslmode=disable")
	db.LogMode(true)
	PanicIf(err)
	DB = &db
	return DB
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
