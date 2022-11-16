package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	_host         = "localhost"
	_user         = "fuever"
	_password     = "fuever"
	_databaseName = "fuever"
	_port         = "5432"
)

var (
	db   *gorm.DB = nil
	_dsn          = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", _host, _user, _password, _databaseName, _port)
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
		&New{},
		&Block{},
	)
	if err != nil {
		panic(err)
	}
	// 查询自增序列 如果已有记录则不设置初值
	rows, err := db.Raw("SELECT last_value from admins_id_seq;").Rows()
	rows.Next()
	lastValue := 1
	rows.Scan(&lastValue)
	// 没有记录就把自增序列初始值置为2000000000
	// 通过区分id的值来区分普通用户和管理员
	if lastValue == 1 {
		db.Exec("SELECT SETVAL('admins_id_seq', 2000000000, false)")
	}
	if err != nil {
		panic(err)
	}
}
