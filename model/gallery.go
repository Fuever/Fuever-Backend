package model

type Gallery struct {
	ID         int     `gorm:"primaryKey;autoIncrement" json:"id"`
	AuthorID   int     `gorm:"column:author_id;not null" json:"author_id"`
	Title      string  `gorm:"column:title;uniqueIndex;not null" json:"title"`
	Content    string  `gorm:"column:content;text;not null" json:"content"`
	Cover      string  `gorm:"column:cover;not null" json:"cover"`
	CreateTime int64   `gorm:"column:create_time;not null" json:"create_time"`
	PostID     int     `gorm:"column:post_id;not null" json:"post_id"` // 关联的帖子
	PositionX  float64 `gorm:"column:position_x;not null" json:"position_x"`
	PositionY  float64 `gorm:"column:position_y;not null" json:"position_y"`
}

func CreateGallery(gallery *Gallery) error {
	err := db.Create(gallery).Error
	if err != nil {
		return err
	}
	return nil
}

func GetGalleryByID(id int) (*Gallery, error) {
	gallery := &Gallery{ID: id}
	err := db.First(gallery).Error
	if err != nil {
		return nil, err
	}
	return gallery, nil
}

func GetAllGalleries() ([]*Gallery, error) {
	galleries := make([]*Gallery, 0)
	err := db.
		Select(
			"id, " +
				"author_id, " +
				"title, " +
				"create_time, " +
				"position_x, " +
				"position_y, " +
				"cover").
		Find(&galleries).
		Error
	return galleries, err
}

func DeleteGalleryID(id int) error {
	err := db.Delete(&Gallery{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

type GalleryInfo struct {
	Gallery
	AuthorName string `json:"author_name,omitempty"`
}

func GetGalleryInfoByID(id int) (*GalleryInfo, error) {
	gallery, err := GetGalleryByID(id)
	if err != nil {
		return nil, err
	}
	info := &GalleryInfo{
		Gallery:    *gallery,
		AuthorName: "",
	}
	// gallery只能管理员注册
	admin, err := GetAdminByID(gallery.AuthorID)
	if err == nil {
		// 有这条记录
		info.AuthorName = admin.Name
	}
	return info, nil
}
