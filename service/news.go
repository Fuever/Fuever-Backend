package service

import "Fuever/model"

type NewsInfo struct {
	model.News
	AuthorName string `json:"author_name,omitempty"`
}

func GetNews(newsID int) (*NewsInfo, error) {
	news, err := model.GetNewsByID(newsID)
	if err != nil {
		return nil, err
	}
	admin, err := model.GetAdminByID(news.AuthorID)
	info := &NewsInfo{
		News:       *news,
		AuthorName: "",
	}
	if err != nil {
		// 创建者销号了
		info.AuthorID = -1
		return info, nil
	}
	info.AuthorName = admin.Name
	return info, nil
}

func GetNewses(offset int, limit int) ([]*NewsInfo, error) {
	newses, err := model.GetNewsWithOffsetLimit(offset, limit)
	if err != nil {
		return nil, err
	}
	m := make(map[int]*model.Admin, 0)
	newsInfo := make([]*NewsInfo, len(newses))
	for i, news := range newses {
		a, flag := m[news.AuthorID]
		if flag {
			newsInfo[i] = &NewsInfo{
				News:       *news,
				AuthorName: a.Name,
			}
		} else {
			// news的创建者只会是管理员
			admin, err := model.GetAdminByID(news.AuthorID)
			if err != nil {
				// 这条记录的创作者已经销毁帐号了
				newsInfo[i] = &NewsInfo{
					News:       *news,
					AuthorName: "",
				}
			} else {
				// 缓存
				newsInfo[i] = &NewsInfo{
					News:       *news,
					AuthorName: admin.Name,
				}
				m[news.AuthorID] = admin
			}
		}
	}
	return newsInfo, nil
}
