package model

type Gallery struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Title      string
	Author     string
	Content    string
	Cover      string
	CreateTime int64
}
