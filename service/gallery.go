package service

import "Fuever/model"

func GetSpecifyGallery(galleryID int) (*model.GalleryInfo, error) {
	return model.GetGalleryInfoByID(galleryID)
}

func GetAllGalleries() ([]*model.Gallery, error) {
	return model.GetAllGalleries()
}
