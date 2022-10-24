package model

import "time"

type Message struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	AuthorID    int       `gorm:"column:author_id;index;not null"`
	Content     string    `gorm:"text;not null"`
	PostID      int       `gorm:"column:post_id;index;not null"`
	CreatedTime time.Time `gorm:"not null"`
}
