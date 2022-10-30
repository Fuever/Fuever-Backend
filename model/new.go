package model

import (
	"time"
)

type New struct {
	ID         int       `gorm:"primaryKey;autoIIncrement"`
	AuthorID   int       `gorm:"column:author_id;not null;index"`
	Title      string    `gorm:"varchar(128);not null"` // 我觉得这个地方要建一个索引啊
	Content    string    `gorm:"text;not null"`
	CreateTime time.Time `gorm:"not null"`
}

func CreateNew(_new *New) error {
	err := db.Create(_new).Error
	if err != nil {
		return err
	}
	return nil
}

// GetNewByID
// 说明一下为什么没有做一个用Title来查询的方法
// 我希望前端展示的时候能把ID和Title绑起来
// 这样可以少建一个索引
// 而且减少网络开销(其实完全没必要
func GetNewByID(id int) (*New, error) {
	_new := &New{
		ID: id,
	}
	err := db.First(_new).Error
	if err != nil {
		return nil, err
	}
	return _new, nil
}

func GetNewsWithLimit(limit int) ([]*New, error) {
	news := make([]*New, 0)
	err := db.Limit(limit).Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil
}

func GetNewsByAuthorIDWIthLimit(authorID int, limit int) ([]*New, error) {
	news := make([]*New, 0)
	err := db.Limit(limit).Where(&New{AuthorID: authorID}).Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil

}

func UpdateNew(_new *New) error {
	err := db.Where("id = ?", _new.ID).Updates(_new).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteNewByID(id int) error {
	err := db.Delete(&New{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
