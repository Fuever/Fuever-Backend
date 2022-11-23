package service

import "Fuever/model"

func GetAnniversary(annivID int) (*model.AnniversaryInfo, error) {
	return model.GetAnniversaryInfoByID(annivID)
}
func GetAnniversaries(offset int, limit int) ([]*model.AnniversaryInfo, error) {
	return model.GetAnniversariesInfoWithOffsetLimit(offset, limit)
}
