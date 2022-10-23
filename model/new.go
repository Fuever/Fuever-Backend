package model

import "time"

type New struct {
	ID         int `gorm:"primaryKey"`
	AuthorID   int
	Title      string // 我觉得这个地方要建一个索引啊
	Content    string
	CreateTime time.Time
}
