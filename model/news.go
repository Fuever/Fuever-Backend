package model

type News struct {
	ID         int    `gorm:"primaryKey;autoIIncrement"`
	AuthorID   int    `gorm:"column:author_id;not null;index"`
	Title      string `gorm:"varchar(128);not null"` // 我觉得这个地方要建一个索引啊
	Content    string `gorm:"text;not null"`
	CreateTime int64  `gorm:"autoCreateTime"`
	Cover      string `gorm:"varchar(128)"` // 新闻的封面
}

func CreateNews(_new *News) error {
	err := db.Create(_new).Error
	if err != nil {
		return err
	}
	return nil
}

// GetNewsByID
// 说明一下为什么没有做一个用Title来查询的方法
// 我希望前端展示的时候能把ID和Title绑起来
// 这样可以少建一个索引
// 而且减少网络开销(其实完全没必要
func GetNewsByID(id int) (*News, error) {
	_new := &News{
		ID: id,
	}
	err := db.First(_new).Error
	if err != nil {
		return nil, err
	}
	return _new, nil
}

func GetNewsWithOffsetLimit(offset int, limit int) ([]*News, error) {
	news := make([]*News, 0)
	err := db.Offset(offset).Limit(limit).Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil
}

func GetNewsByAuthorIDWIthOffsetLimit(authorID int, offset int, limit int) ([]*News, error) {
	news := make([]*News, 0)
	err := db.Offset(offset).Limit(limit).Where(&News{AuthorID: authorID}).Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil

}

func UpdateNew(_new *News) error {
	err := db.Omit("ID").Where("id = ?", _new.ID).Updates(_new).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteNewByID(id int) error {
	err := db.Delete(&News{}, id).Error
	if err != nil {
		return err
	}
	return nil
}