package model

import "time"

type Anniversary struct {
	ID       int       `gorm:"primaryKey;autoIIncrement"`
	AdminID  int       `gorm:"column:admin_id;index;not null"`
	Title    string    `gorm:"varchar(128);not null"`
	Content  string    `gorm:"text:not null"`
	Start    time.Time `gorm:"not null"`
	End      time.Time `gorm:"not null"`
	Location string    `gorm:"varchar(128);not null"`
}
