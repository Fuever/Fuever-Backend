package model

type Block struct {
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"varchar(128);not null;uniqueIndex"`
}

func CreateBlock(block *Block) error {
	err := db.Create(block).Error
	if err != nil {
		return err
	}
	return nil
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
	err := db.Where("id = ?", block.ID).Updates(block).Error
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
