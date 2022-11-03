package model

import (
	"time"
)

type Message struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	AuthorID    int       `gorm:"column:author_id;index;not null"`
	Content     string    `gorm:"text;not null"`
	PostID      int       `gorm:"column:post_id;index;not null"`
	CreatedTime time.Time `gorm:"not null"`
}

func CreateMessage(message *Message) error {
	err := db.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}

//func GetAllMessageByPostID(postID int) ([]*Message, error) {
//	messages := make([]*Message, 0)
//	err := db.Where("post_id = ?", postID).Find(&messages).Error
//	if err != nil {
//		return nil, err
//	}
//	return messages, nil
//}

func GetMessageByPostIDWithOffsetLimit(postID int, offset int, limit int) ([]*Message, error) {
	messages := make([]*Message, 0)
	err := db.Where("post_id = ?", postID).Offset(offset).Limit(limit).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetMessageByAuthorIDWithOffsetLimit(authorID int, offset int, limit int) ([]*Message, error) {
	messages := make([]*Message, 0)
	err := db.Where("author_id = ?", authorID).Offset(offset).Limit(limit).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func DeleteMessageByID(id int) error {
	err := db.Delete(&Message{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
