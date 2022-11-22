package model

type Anniversary struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	AdminID   int     `gorm:"column:admin_id;index;not null"`
	Title     string  `gorm:"varchar(128);not null"`
	Content   string  `gorm:"text:not null"`
	Start     int64   `gorm:"not null"`
	End       int64   `gorm:"not null"`
	PositionX float64 `gorm:"position_x;not null"`
	PositionY float64 `gorm:"position_y;not null"`
}

func CreateAnniversary(anniversary *Anniversary) error {
	err := db.Create(anniversary).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAnniversaryByID(id int) (*Anniversary, error) {
	anniv := &Anniversary{
		ID: id,
	}
	err := db.First(anniv).Error
	if err != nil {
		return nil, err
	}
	return anniv, err
}

func GetAnniversariesWithOffsetLimit(offset int, limit int) ([]*Anniversary, error) {
	anniversaries := make([]*Anniversary, 0)
	err := db.Select("ID", "AdminID", "Title", "Cover").Limit(limit).Offset(offset).Find(&anniversaries).Error
	if err != nil {
		return nil, err
	}
	return anniversaries, nil
}

func GetAnniversariesByAdminIDWIthOffsetLimit(adminID int, offset int, limit int) ([]*Anniversary, error) {
	anniversaries := make([]*Anniversary, 0)
	err := db.Select("ID", "AdminID", "Title", "Cover").Limit(limit).Offset(offset).Where(&Anniversary{AdminID: adminID}).Find(&anniversaries).Error
	if err != nil {
		return nil, err
	}
	return anniversaries, nil

}

func UpdateAnniversaryByID(anniversary *Anniversary) error {
	err := db.Omit("ID").Where("id = ?", anniversary.ID).Updates(anniversary).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAnniversaryByID(id int) error {
	err := db.Delete(&Anniversary{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
