package service

import "Fuever/model"

func GetNews(newsID int) (*model.NewsInfo, error) {
	return model.GetNewsInfo(newsID)
}

func GetNewses(offset int, limit int) ([]*model.NewsInfo, error) {
	return model.GetNewsesInfo(offset, limit)
}
