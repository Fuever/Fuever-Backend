package model

// Message
// Message只允许增删 不允许修改
type Message struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	AuthorID    int    `gorm:"column:author_id;index;not null" json:"author_id"`
	Content     string `gorm:"text;not null" json:"content"`
	PostID      int    `gorm:"column:post_id;index;not null" json:"post_id"`
	CreatedTime int64  `gorm:"column:created_time" json:"created_time"`
}

func CreateMessage(message *Message) error {
	err := db.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}

//func GetAllMessageByPostID(postID int) ([]*Message, error) {
//	messages := make([]*Message, 0)
//	err := db.Where("post_id = ?", postID).Find(&messages).Error
//	if err != nil {
//		return nil, err
//	}
//	return messages, nil
//}

func GetMessageByID(id int) (*Message, error) {
	msg := &Message{ID: id}
	err := db.Find(msg).Error
	if err != nil {
		return nil, err
	}
	return msg, err
}

func GetMessageByPostIDWithOffsetLimit(postID int, offset int, limit int) ([]*Message, error) {
	messages := make([]*Message, 0)
	err := db.Where("post_id = ?", postID).Offset(offset).Limit(limit).Order("created_time").Find(&messages).Error
	return messages, err
}

func GetMessageByAuthorIDWithOffsetLimit(authorID int, offset int, limit int) ([]*Message, error) {
	messages := make([]*Message, 0)
	err := db.Where("author_id = ?", authorID).Offset(offset).Limit(limit).Find(&messages).Error
	return messages, err
}

func DeleteMessageByID(id int) error {
	err := db.Delete(&Message{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
