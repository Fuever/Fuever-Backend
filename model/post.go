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

func GetNormalPostsWithLimit(limit int) ([]*Post, error) {
	return getParticularStatePostWithLimit(normal, limit)
}

func GetHidePostsWithLimit(limit int) ([]*Post, error) {
	return getParticularStatePostWithLimit(hide, limit)
}

func GetTopPostsWithLimit(limit int) ([]*Post, error) {
	return getParticularStatePostWithLimit(top, limit)
}

func GetPostsByAuthorIDWIthLimit(authorID int, limit int) ([]*Post, error) {
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

func getParticularStatePostWithLimit(state int, limit int) ([]*Post, error) {
	posts := make([]*Post, 0)
	err := db.Where(&Post{State: state}).Limit(limit).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
