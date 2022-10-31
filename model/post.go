package model

import (
	"time"
)

type Post struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	AuthorID    int       `gorm:"column:author_id;index;not null"`
	Title       string    `gorm:"varchar(128);not null"`
	CreatedTime time.Time `gorm:"not null"`
	UpdatedTime time.Time `gorm:"not null"`
	State       int       `gorm:"not null"`
}

const (
	normal = iota // 常态
	hide          // 隐藏
	top           // 置顶
)

func CreatePost(post *Post) error {
	err := db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func GetNormalPostsWithOffsetLimit(offset int, limit int) ([]*Post, error) {
	return getParticularStatePostWithOffsetLimit(normal, offset, limit)
}

func GetHidePostsWithOffsetLimit(offset int, limit int) ([]*Post, error) {
	return getParticularStatePostWithOffsetLimit(hide, offset, limit)
}

func GetTopPostsWithOffsetLimit(offset int, limit int) ([]*Post, error) {
	return getParticularStatePostWithOffsetLimit(top, offset, limit)
}

func GetPostsByAuthorIDWIthOffsetLimit(authorID int, limit int) ([]*Post, error) {
	posts := make([]*Post, 0)
	err := db.Limit(limit).Where(&Post{AuthorID: authorID}).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil

}

func UpdatePost(post *Post) error {
	err := db.Where("id = ?", post.ID).Updates(post).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePostByID(id int) error {
	err := db.Delete(&Post{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func getParticularStatePostWithOffsetLimit(state int, offset int, limit int) ([]*Post, error) {
	posts := make([]*Post, 0)
	err := db.Where(&Post{State: state}).Offset(offset).Limit(limit).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
