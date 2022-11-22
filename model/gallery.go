package model

type Gallery struct {
	ID         int     `gorm:"primaryKey;autoIncrement"`
	AuthorID   int     `gorm:"author_id;not null"`
	Title      string  `gorm:"title;not null"`
	Content    string  `gorm:"text;not null"`
	Cover      string  `gorm:"cover;not null"`
	CreateTime int64   `gorm:"create_time;not null"`
	PostID     int     `gorm:"post_id;not null"`
	PositionX  float64 `gorm:"position_x;not null"`
	PositionY  float64 `gorm:"position_y;not null"`
}
