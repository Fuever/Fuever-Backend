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
	posts := make([]*Post, 0)
	err := db.Where(&Post{State: state}).Limit(limit).Find(&posts).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return posts
}

func GetPostsByAuthorIDWIthLimit(authorID int, limit int) []*Post {
	posts := make([]*Post, 0)
	err := db.Limit(limit).Where(&Post{AuthorID: authorID}).Find(&posts).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return posts

}

func UpdatePostByID(post *Post) {
	err := db.Where("id = ?", post.ID).Updates(post).Error
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
