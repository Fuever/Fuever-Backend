package model

type Block struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title    string `gorm:"varchar(128);not null;uniqueIndex" json:"title"`
	AuthorID int    `gorm:"column:author_id;not null;index" json:"authorID"`
}

func CreateBlock(block *Block) error {
	err := db.Create(block).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBlockWithLimitOffset(limit int, offset int) ([]*Block, error) {
	blocks := make([]*Block, 0)
	err := db.Offset(offset).Limit(limit).Find(&blocks).Error
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func GetBlockByID(id int) (*Block, error) {
	block := &Block{ID: id}
	err := db.First(block).Error
	if err != nil {
		return nil, err
	}
	return block, nil
}

func UpdateBlock(block *Block) error {
	err := db.Omit("ID").Where("id = ?", block.ID).Updates(block).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteBlockByID(id int) error {
	err := db.Delete(&Block{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}

type BlockInfo struct {
	Block
	AuthorName string `json:"author_name,omitempty"`
}

func GetBlockAndAuthorNameWithLimitOffset(limit int, offset int) ([]*BlockInfo, error) {
	info := make([]*BlockInfo, 0)
	err := db.Model(&Block{}).
		Select("blocks.id, title, author_id, name as author_name").
		Joins("join  admins  on blocks.author_id = admins.id").
		Offset(offset).
		Limit(limit).
		Scan(&info).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}
