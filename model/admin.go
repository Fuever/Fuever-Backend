package model

type Admin struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"varchar(128);not null"`
	Password string `gorm:"varchar(64);not null"`
}
