package model

type News struct {
	ID         int    `gorm:"primaryKey;autoIIncrement" json:"id"`
	AuthorID   int    `gorm:"column:author_id;not null;index" json:"author_id"`
	Title      string `gorm:"varchar(128);not null; index" json:"title"` // 我觉得这个地方要建一个索引啊
	Content    string `gorm:"text;not null" json:"content"`
	CreateTime int64  `gorm:"column:created_time" json:"create_time"`
	Cover      string `gorm:"varchar(128)" json:"cover"` // 新闻的封面
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

func GetNewsesWithOffsetLimit(offset int, limit int) ([]*News, error) {
	news := make([]*News, 0)
	err := db.Select("id", "author_id", "title", "created_time", "cover").Offset(offset).Limit(limit).Find(&news).Error
	return news, err
}

func GetNewsesByAuthorIDWIthOffsetLimit(authorID int, offset int, limit int) ([]*News, error) {
	news := make([]*News, 0)
	err := db.Offset(offset).Limit(limit).Where(&News{AuthorID: authorID}).Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil

}

func UpdateNews(_new *News) error {
	err := db.Omit("ID").Where("id = ?", _new.ID).Updates(_new).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteNewsByID(id int) error {
	err := db.Delete(&News{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

type NewsInfo struct {
	News
	AuthorName string `json:"author_name,omitempty"`
}

func GetNewsInfo(newsID int) (*NewsInfo, error) {
	news, err := GetNewsByID(newsID)
	if err != nil {
		return nil, err
	}
	admin, err := GetAdminByID(news.AuthorID)
	info := &NewsInfo{
		News:       *news,
		AuthorName: "",
	}
	if err != nil {
		// 创建者销号了
		return info, nil
	}
	info.AuthorName = admin.Name
	return info, nil
}

func GetNewsesInfo(offset int, limit int) ([]*NewsInfo, error) {
	info := make([]*NewsInfo, 0)
	err := db.Model(&News{}).
		Select("news.id, news.author_id, content, title, created_time, cover, admins.name as author_name").
		Joins("join admins on news.author_id=admins.id").
		Offset(offset).
		Limit(limit).
		Scan(&info).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(info); i++ {
		s := info[i].Content
		if len(s) > 20 {
			info[i].Content = s[:20]
		}
	}
	return info, nil
}
