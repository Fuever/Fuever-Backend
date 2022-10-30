package model

type Admin struct {
	ID       int    `gorm:"primaryKey;autoIncrement;check:id > 2000000000"`
	Username string `gorm:"varchar(128);not null"`
	Password string `gorm:"varchar(64);not null"`
}

func CreateAdmin(admin *Admin) error {
	err := db.Create(admin).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAdminByID(id int) (*Admin, error) {
	admin := &Admin{ID: id}
	err := db.Find(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func UpdateAdmin(admin *Admin) error {
	err := db.Where("id = ?", admin.ID).Updates(admin).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAdminByID(id int) error {
	err := db.Delete(&Admin{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
