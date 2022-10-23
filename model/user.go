package model

import "time"

type User struct {
	ID           int    `gorm:"primaryKey"`
	Mail         string `gorm:"unique"`
	Username     string `gorm:"unique"`
	Password     string
	Nickname     string `gorm:"index"`
	Avatar       string
	Phone        int `gorm:"unique"`
	Gender       bool
	Age          int
	Job          string
	EntranceTime time.Time
	ClassID      int
}
