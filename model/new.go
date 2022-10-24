package model

import "time"

type New struct {
	ID         int       `gorm:"primaryKey;autoIIncrement"`
	AuthorID   int       `gorm:"column:author_id;not null"`
	Title      string    `gorm:"varchar(128);not null"` // 我觉得这个地方要建一个索引啊
	Content    string    `gorm:"text;not null"`
	CreateTime time.Time `gorm:"not null"`
}
