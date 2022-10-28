package model

import (
	"log"
	"time"
)

type New struct {
	ID         int       `gorm:"primaryKey;autoIIncrement"`
	AuthorID   int       `gorm:"column:author_id;not null;index"`
	Title      string    `gorm:"varchar(128);not null"` // 我觉得这个地方要建一个索引啊
	Content    string    `gorm:"text;not null"`
	CreateTime time.Time `gorm:"not null"`
}

// GetNewByID
// 说明一下为什么没有做一个用Title来查询的方法
// 我希望前端展示的时候能把ID和Title绑起来
// 这样可以少建一个索引
// 而且减少网络开销(其实完全没必要
func GetNewByID(id int) *New {
	_new := &New{
		ID: id,
	}
	err := db.First(_new).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return _new
}

func GetNewsWithLimit(limit int) []*New {
	news := make([]*New, 0)
	err := db.Limit(limit).Find(&news).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return news
}

func GetNEwsByAuthorIDWIthLimit(authorID int, limit int) []*New {
	news := make([]*New, 0)
	err := db.Limit(limit).Where(&New{AuthorID: authorID}).Find(&news).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	return news

}
