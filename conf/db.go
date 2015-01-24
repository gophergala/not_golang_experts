package conf

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupDB() gorm.DB {
	db, err := gorm.Open("postgres", "dbname=gopherstalker sslmode=disable")
	db.LogMode(true)
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
