package model

type Anniversary struct {
	ID        int     `gorm:"primaryKey;autoIncrement" json:"id"`
	AdminID   int     `gorm:"column:admin_id;index;not null" json:"admin_id"`
	Title     string  `gorm:"varchar(128);not null" json:"title"`
	Content   string  `gorm:"text:not null" json:"content"`
	Start     int64   `gorm:"not null" json:"start"`
	End       int64   `gorm:"not null" json:"end"`
	PositionX float64 `gorm:"position_x;not null" json:"position_x"`
	PositionY float64 `gorm:"position_y;not null" json:"position_y"`
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

type AnniversaryInfo struct {
	Anniversary
	AuthorName string `json:"author_name,omitempty"`
}

func GetAnniversaryInfoByID(id int) (*AnniversaryInfo, error) {
	anniv, err := GetAnniversaryByID(id)
	if err != nil {
		return nil, err
	}
	authorName := ""
	admin, err := GetAdminByID(anniv.AdminID)
	if err == nil {
		authorName = admin.Name
	}
	info := &AnniversaryInfo{
		Anniversary: *anniv,
		AuthorName:  authorName,
	}
	return info, nil
}
func GetAnniversariesInfoWithOffsetLimit(offset int, limit int) ([]*AnniversaryInfo, error) {
	annivInfo := make([]*AnniversaryInfo, 0)
	err := db.Model(&Anniversary{}).Select("anniversaries.id, " +
		"anniversaries.admin_id, " +
		"anniversaries.title, " + // 不返回content?
		"anniversaries.start, " +
		"anniversaries.end," +
		"anniversaries.position_x, " +
		"anniversaries.position_y, " +
		"admins.name").
		Joins("join admins on admins.id=anniversaries.admin_id").
		Offset(offset).
		Limit(limit).
		Scan(&annivInfo).Error
	return annivInfo, err
}
