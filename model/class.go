package model

import "errors"

type Class struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	ClassName string `gorm:"class_name;index;not null"`
	StudentID int    `gorm:"student_id;index;not null"`
}

func CreateClass(class *Class) error {
	cnt := int64(0)
	db.Model(&Class{}).
		Where("class_name = ? and student_id = ?", class.ClassName, class.StudentID).
		Count(&cnt)
	if cnt == 0 {
		err := db.Create(class).Error
		if err != nil {
			return err
		}
		return nil
		/// 没有这一组关系
	} else {
		return errors.New("duplicate")
	}
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
		Select("users.id, users.username, users.student_id").
		Joins("join classes on classes.student_id = users.student_id").
		Where("classes.class_name = ?", className).
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

func GetClassesByStudentID(studentID int) ([]*Class, error) {
	res := make([]*Class, 0)
	err := db.Model(&Class{}).
		Where("student_id = ?", studentID).
		Scan(&res).
		Error
	return res, err
}

func GetClassesByFuzzyQuery(word string) ([]*Class, error) {
	res := make([]*Class, 0)
	err := db.Model(&Class{}).
		Select("class_name").
		Group("class_name").
		Having("class_name like ?", "%"+word+"%").
		Scan(&res).
		Error
	return res, err
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
