package model

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement;check: id < 2000000000"`
	Mail         string `gorm:"uniqueIndex;varchar(64);not null"`
	Password     string `gorm:"varchar(64);not null"`
	Username     string `gorm:"index;varchar(32);"` // 真名需要验证
	Nickname     string `gorm:"index;varchar(32)"`
	Avatar       string `gorm:"varchar(64)"`
	StudentID    int    `gorm:"column:student_id;index"` //未验证时为空
	Phone        int    `gorm:"index"`
	Gender       bool
	Age          int
	Job          string
	EntranceTime int64
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
	err := db.First(user).Error
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

func IsIDBelongToUser(id int) bool {
	return id < 2_000_000_000
}

func GetUserWithOffsetLimit(offset int, limit int) ([]*User, error) {
	users := make([]*User, 0)
	err := db.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func PickRandomUser(ignoreUserID int, limit int) ([]*User, error) {
	users := make([]*User, 0)
	err := db.Raw(
		"SELECT * FROM users WHERE (users.id != ?) AND (users.student_id != 0) AND (users.id >= RANDOM() * (SELECT MAX(users.id) FROM users)) LIMIT ?",
		ignoreUserID, limit).Scan(&users).Error
	return users, err
}
