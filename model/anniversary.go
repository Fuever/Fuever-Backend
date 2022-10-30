package model

import "time"

type Anniversary struct {
	ID       int       `gorm:"primaryKey;autoIncrement"`
	AdminID  int       `gorm:"column:admin_id;index;not null"`
	Title    string    `gorm:"varchar(128);not null"`
	Content  string    `gorm:"text:not null"`
	Start    time.Time `gorm:"not null"`
	End      time.Time `gorm:"not null"`
	Location string    `gorm:"varchar(128);not null"`
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

func GetAnniversariesWithLimit(limit int) ([]*Anniversary, error) {
	anniversaries := make([]*Anniversary, 0)
	err := db.Limit(limit).Find(&anniversaries).Error
	if err != nil {
		return nil, err
	}
	return anniversaries, nil
}

func GetAnniversariesByAdminIDWIthLimit(adminID int, limit int) ([]*Anniversary, error) {
	anniversaries := make([]*Anniversary, 0)
	err := db.Limit(limit).Where(&Anniversary{AdminID: adminID}).Find(&anniversaries).Error
	if err != nil {
		return nil, err
	}
	return anniversaries, nil

}

func UpdateAnniversaryByID(anniversary *Anniversary) error {
	err := db.Where("id = ?", anniversary.ID).Updates(anniversary).Error
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
