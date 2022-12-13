package model

type Class struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	ClassName string `gorm:"class_name;index;not null"`
	StudentID int    `gorm:"student_id;index;not null"`
}

func CreateClass(class *Class) error {
	err := db.Create(class).Error
	if err != nil {
		return err
	}
	return nil
}

func GetClassList(offset, limit int) ([]*Class, error) {
	classes := make([]*Class, 0)
	err := db.Model(&Class{}).
		Offset(offset).
		Limit(limit).
		Group("class_name").
		Select("class_name").
		Scan(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, err
}

// GetStudentListByClassName
// 这个方法返回的列表中 用户仅仅具有id username student_id 三个属性
func GetStudentListByClassName(className string) ([]*User, error) {
	users := make([]*User, 0)
	err := db.Model(&User{}).
		Select("user.id, user.username, user.student_id").
		Joins("join class on class.student_id = user.id").
		Where("class.class_name = ?", className).
		Scan(&users).Error
	return users, err
}

func CountStudentJoinedClassNumber(studentID int) (int64, error) {
	cnt := int64(0)
	err := db.Model(&Class{}).
		Where("student_id = ?", studentID).
		Count(&cnt).
		Error
	return cnt, err
}

func GetClassByID(id int) (*Class, error) {
	class := &Class{
		ID: id,
	}
	err := db.First(class).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}

func DeleteClassByID(id int) error {
	err := db.Delete(&Class{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
