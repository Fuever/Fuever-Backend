package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host         = "localhost"
	user         = "gorm"
	password     = "gorm"
	databaseName = "gorm"
	port         = "9910"
)

var (
	db  *gorm.DB = nil
	dsn          = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, databaseName, port)
)

func InitDB() {
	dbEngine, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// can't connect database
		log.Fatalln(err)
	}
	db = dbEngine
	// automatically generate database tables
	err = db.AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
