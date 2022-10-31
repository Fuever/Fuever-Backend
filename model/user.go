package model

import "time"

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement;check: id < 2000000000"`
	Mail         string `gorm:"uniqueIndex;varchar(64);not null"`
	Username     string `gorm:"index;varchar(32);not null"`
	Password     string `gorm:"varchar(64);not null"`
	Nickname     string `gorm:"index;varchar(32)"`
	Avatar       string `gorm:"varchar(64)"`
	StudentID    int    `gorm:"column:student_id;index"` // 这地方加索引好像会比较好 但是问题在于不应该为空
	Phone        int    `gorm:"uniqueIndex"`
	Gender       bool
	Age          int
	Job          string
	EntranceTime time.Time
	ClassID      int `gorm:"column:class_id;index"`
	Residence    string
}

// CreateUser
// 调用这个方法之后
// id会被自动注入到传入的结构体里
// 什么 你说这个不符合不可变性？
// 要是可以我也不想用这种东西啊
func CreateUser(user *User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(id int) (*User, error) {
	user := &User{ID: id}
	err := db.First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsersByStudentID(studentID int) ([]*User, error) {
	user := make([]*User, 0)
	err := db.Where("student_id = ?", studentID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *User) error {
	err := db.Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByID(id int) error {
	err := db.Delete(&User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
