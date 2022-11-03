package model

type Class struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	ClassName string `gorm:"class_name;uniqueIndex;not null"`
	Major     string `gorm:"varchar(128);index;not null"`
	Grade     int    `gorm:"not null"`
}

func CreateClass(class *Class) error {
	err := db.Create(class).Error
	if err != nil {
		return err
	}
	return nil
}

func GetClassByMajor(major string) (*Class, error) {
	class := &Class{}
	err := db.Where("major = ?", major).First(&class).Error
	if err != nil {
		return nil, err
	}
	return class, err
}

func GetClassByGrade(grade int) ([]*Class, error) {
	classes := make([]*Class, 0)
	err := db.Where("grade = ?", grade).Find(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func GetClassesByClassNameFuzzyQuery(fuzzyClassName string) ([]*Class, error) {
	classes := make([]*Class, 0)
	err := db.Where("class_name like ?", "%"+fuzzyClassName+"%").Find(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, err
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
