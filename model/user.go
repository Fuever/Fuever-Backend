package model

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement;check: id < 2000000000"`
	Mail         string `gorm:"uniqueIndex;varchar(64);not null"`
	Password     string `gorm:"varchar(64);not null"`
	Username     string `gorm:"index;varchar(32);"` // 真名需要验证
	Nickname     string `gorm:"index;varchar(32)"`
	Avatar       string `gorm:"varchar(64)"`
	StudentID    int    `gorm:"column:student_id;uniqueIndex"` // 这地方加索引好像会比较好 但是问题在于不应该为空
	Phone        int    `gorm:"uniqueIndex"`
	Gender       bool
	Age          int
	Job          string
	EntranceTime int64
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

func GetUserByMailbox(mailbox string) (*User, error) {
	user := &User{}
	err := db.Where("mail", mailbox).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(id int) (*User, error) {
	user := &User{ID: id}
	err := db.Select("ID", "Password").First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsersByStudentID(studentID int) ([]*User, error) {
	user := make([]*User, 0)
	err := db.Select("ID", "Nickname", "Avatar", "StudentID", "Gender").Where("student_id = ?", studentID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *User) error {
	err := db.Omit("ID").Where("id = ?", user.ID).Updates(user).Error
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
