package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host         = "localhost"
	user         = "fuever"
	password     = "fuever"
	databaseName = "fuever"
	port         = "5432"
)

var (
	db   *gorm.DB = nil
	_dsn          = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, databaseName, port)
)

func InitDB() {
	dbEngine, err := gorm.Open(postgres.Open(_dsn), &gorm.Config{})
	if err != nil {
		// can't connect database
		log.Fatalln(err)
	}
	db = dbEngine
	// automatically generate database tables
	err = db.AutoMigrate(
		&User{},
		&Admin{},
		&Anniversary{},
		&Class{},
		&Message{},
		&Post{},
		&New{})
	if err != nil {
		log.Fatalln(err)
	}
}
