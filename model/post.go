package model

type Post struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	AuthorID    int    `gorm:"column:author_id;index;not null" json:"author_id"`
	Title       string `gorm:"varchar(128);index;not null" json:"title"`
	CreatedTime int64  `gorm:"column:created_time;index" json:"created_time"`
	UpdatedTime int64  `gorm:"column:updated_time;not null;index" json:"updated_time"`
	State       int    `gorm:"not null" json:"state"`
	BlockID     int    `gorm:"column:block_id;not null" json:"block_id"`
	IsLock      bool   `gorm:"column:is_lock;default:0" json:"is_lock"`
}

const (
	_normal = iota // 常态
	_hide          // 隐藏
	_top           // 置顶
)

func CreatePost(post *Post) error {
	err := db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPostByID(id int) (*Post, error) {
	post := &Post{
		ID: id,
	}
	err := db.First(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetAllNormalPostsWithOffsetLimit(offset int, limit int) ([]*Post, error) {
	return getParticularStatePostsWithOffsetLimit(_normal, offset, limit)
}

func GetNormalPostsWithOffsetLimit(blockID int, offset int, limit int) ([]*Post, error) {
	return getParticularStatePostWithOffsetLimit(blockID, _normal, offset, limit)
}

func GetHidePostsWithOffsetLimit(blockID int, offset int, limit int) ([]*Post, error) {
	return getParticularStatePostWithOffsetLimit(blockID, _hide, offset, limit)
}

func GetTopPostsWithOffsetLimit(blockID int, offset int, limit int) ([]*Post, error) {
	return getParticularStatePostWithOffsetLimit(blockID, _top, offset, limit)
}

func GetPostsByAuthorIDWIthOffsetLimit(authorID int, limit int) ([]*Post, error) {
	posts := make([]*Post, 0)
	err := db.Limit(limit).Where(&Post{AuthorID: authorID}).Find(&posts).Error
	return posts, err

}

func UpdatePost(post *Post) error {
	err := db.Omit("ID").Where("id = ?", post.ID).Updates(post).Error
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

func getParticularStatePostWithOffsetLimit(blockID int, state int, offset int, limit int) ([]*Post, error) {
	posts := make([]*Post, 0)
	err := db.Where(&Post{State: state, BlockID: blockID}).Offset(offset).Limit(limit).Order("updated_time desc").Find(&posts).Error
	return posts, err
}

func getParticularStatePostsWithOffsetLimit(state int, offset int, limit int) ([]*Post, error) {
	posts := make([]*Post, 0)
	err := db.Where(&Post{State: state}).Offset(offset).Limit(limit).Order("updated_time desc").Find(&posts).Error
	return posts, err
}
