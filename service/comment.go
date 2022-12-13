package service

import "Fuever/model"

type CommentInfo struct {
	model.Message
	AuthorName   string `json:"author_name,omitempty"`
	AuthorAvatar string `json:"author_avatar,omitempty"`
}

func GetComments(postID int, offset int, limit int) ([]*CommentInfo, error) {
	messages, err := model.GetMessageByPostIDWithOffsetLimit(postID, offset, limit)
	if err != nil {
		return nil, err
	}
	comments := make([]*CommentInfo, len(messages))
	// authorID map to user
	m := make(map[int]*model.User)
	for i, message := range messages {
		user, flag := m[message.AuthorID]
		authorName := ""
		authorAvatar := ""
		if flag {
			authorName = user.Nickname
			authorAvatar = user.Avatar
		} else {
			user, err := model.GetUserByID(message.AuthorID)
			if err == nil {
				m[message.AuthorID] = user
				authorAvatar = user.Avatar
				authorName = user.Nickname
			} else {
				// 如果查不到这个信息 就说明发评论的销号了
				m[message.AuthorID] = &model.User{
					ID:           0,
					Mail:         "",
					Password:     "",
					Username:     "",
					Nickname:     "",
					Avatar:       "",
					StudentID:    0,
					Phone:        0,
					Gender:       false,
					Age:          0,
					Job:          "",
					EntranceTime: 0,
					Residence:    "",
				}
				// 填空值作缓存
				// 避免每次销号都重新查一遍表
			}
		}
		comments[i] = &CommentInfo{
			Message:      *message,
			AuthorName:   authorName,
			AuthorAvatar: authorAvatar,
		}
	}
	return comments, nil

}
