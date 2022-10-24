package model

import "time"

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Mail         string `gorm:"uniqueIndex;varchar(64);not null"`
	Username     string `gorm:"index;varchar(32);not null"`
	Password     string `gorm:"varchar(64);not null"`
	Nickname     string `gorm:"index;varchar(32)"`
	Avatar       string `gorm:"varchar(64)"`
	Phone        int    `gorm:"uniqueIndex"`
	Gender       bool
	Age          int
	Job          string
	EntranceTime time.Time
	ClassID      int `gorm:"column:class_id;index"`
}
