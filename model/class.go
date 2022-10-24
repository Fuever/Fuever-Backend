package model

type Class struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	ClassName string `gorm:"class_name;uniqueIndex;not null"`
	major     string `gorm:"varchar(128);index;not null"`
	grade     int    `gorm:"not null"`
}
