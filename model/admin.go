package model

type Admin struct {
	ID       int    `gorm:"primaryKey;autoIncrement;check:id > 2000000000"`
	Username string `gorm:"varchar(128);not null"`
	Password string `gorm:"varchar(64);not null"`
}

func GetAdminByID(id int) (*Admin, error) {
	admin := &Admin{ID: id}
	err := db.Find(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}
