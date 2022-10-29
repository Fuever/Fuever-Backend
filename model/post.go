package model

import (
	"log"
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

func GetNormalPostsWithLimit(limit int) []*Post {
	return getParticularStatePostWithLimit(normal, limit)
}

func GetHidePostsWithLimit(limit int) []*Post {
	return getParticularStatePostWithLimit(hide, limit)
}

func GetTopPostsWithLimit(limit int) []*Post {
	return getParticularStatePostWithLimit(top, limit)
}

func getParticularStatePostWithLimit(state int, limit int) []*Post {
	Posts := make([]*Post, 0)
	err := db.Where(&Post{State: state}).Limit(limit).Find(&Posts).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return Posts
}

func GetPostsByAuthorIDWIthLimit(authorID int, limit int) []*Post {
	Posts := make([]*Post, 0)
	err := db.Limit(limit).Where(&Post{AuthorID: authorID}).Find(&Posts).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return Posts

}

func UpdatePostByID(_Post *Post) {
	err := db.Where("id = ?", _Post.ID).Updates(_Post).Error
	if err != nil {
		log.Println(err)
	}
}

func DeletePostByID(id int) {
	err := db.Delete(&Post{}, id).Error
	if err != nil {
		log.Println(err)
		return
	}
}
