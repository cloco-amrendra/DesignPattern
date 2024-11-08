package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once // This is a struct that will help us to run the initialization function only once.
)

func GetDB() *gorm.DB {
	once.Do(func() { // This function will run only once.
		dsn := "user:password@tcp(db:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		db = database
	})
	return db
}
