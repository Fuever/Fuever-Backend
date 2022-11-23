package service

import "Fuever/model"

func GetNews(newsID int) (*model.NewsInfo, error) {
	return model.GetNewsInfo(newsID)
}

func GetNewses(offset int, limit int) ([]*model.NewsInfo, error) {
	return model.GetNewsesInfo(offset, limit)
	//newses, err := model.GetNewsWithOffsetLimit(offset, limit)
	//if err != nil {
	//	return nil, err
	//}
	//m := make(map[int]*model.Admin, 0)
	//newsInfo := make([]*NewsInfo, len(newses))
	//for i, news := range newses {
	//	a, flag := m[news.AuthorID]
	//	if flag {
	//		newsInfo[i] = &NewsInfo{
	//			News:       *news,
	//			AuthorName: a.Name,
	//		}
	//	} else {
	//		// news的创建者只会是管理员
	//		admin, err := model.GetAdminByID(news.AuthorID)
	//		if err != nil {
	//			// 这条记录的创作者已经销毁帐号了
	//			newsInfo[i] = &NewsInfo{
	//				News:       *news,
	//				AuthorName: "",
	//			}
	//		} else {
	//			// 缓存
	//			newsInfo[i] = &NewsInfo{
	//				News:       *news,
	//				AuthorName: admin.Name,
	//			}
	//			m[news.AuthorID] = admin
	//		}
	//	}
	//}
	//return newsInfo, nil
}
