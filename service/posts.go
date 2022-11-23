package service

import "Fuever/model"

type PostInfo struct {
	model.Post
	AuthorName   string `json:"author_name,omitempty"`
	AuthorAvatar string `json:"author_avatar,omitempty"`
}

func GetPost(postID int) (*PostInfo, error) {
	post, err := model.GetPostByID(postID)
	if err != nil {
		// err not record be found
		return nil, err
	}
	info := &PostInfo{
		Post:         *post,
		AuthorName:   "",
		AuthorAvatar: "",
	}
	if model.IsIDBelongToAdmin(info.AuthorID) {
		// 发帖人是管理员
		admin, err := model.GetAdminByID(info.AuthorID)
		if err != nil {
			// 管理员已注销
			info.AuthorID = -1
			return info, nil
		}
		// 管理员没头像啊...
		info.AuthorName = admin.Name
		//info.AuthorAvatar = DefaultAvatar //TODO 也许管理员需要一个默认的头像
		return info, nil
	} else {
		// 发帖人是用户
		user, err := model.GetUserByID(info.ID)
		if err != nil {
			// 用户已注销
			info.AuthorID = -1
			return info, nil
		}
		info.AuthorName = user.Nickname
		info.AuthorAvatar = user.Avatar
		return info, nil
	}
}

func GetPosts(blockID int, offset int, limit int) ([]*PostInfo, error) {
	posts, err := model.GetNormalPostsWithOffsetLimit(blockID, offset, limit)
	if err != nil {
		return nil, err
	}
	// author id map to admin/user
	m := make(map[int]struct {
		AuthorName   string
		AuthorAvatar string
	}, 0)
	postInfo := make([]*PostInfo, len(posts))
	for i, post := range posts {
		info := &PostInfo{
			Post:         *post,
			AuthorName:   "",
			AuthorAvatar: "",
		}
		temp, flag := m[post.AuthorID]
		if flag {
			// 不需要查数据库
			info.AuthorName = temp.AuthorName
			info.AuthorAvatar = temp.AuthorAvatar
		} else {
			if model.IsIDBelongToAdmin(info.AuthorID) {
				// 发帖人是管理员
				admin, err := model.GetAdminByID(info.AuthorID)
				if err != nil {
					// 管理员已注销
					m[post.AuthorID] = struct {
						AuthorName   string
						AuthorAvatar string
					}{
						AuthorAvatar: "",
						AuthorName:   "",
					}
					continue
				}
				// 管理员没头像啊...
				info.AuthorName = admin.Name
				//info.AuthorAvatar = DefaultAvatar //TODO 也许管理员需要一个默认的头像
				m[post.AuthorID] = struct {
					AuthorName   string
					AuthorAvatar string
				}{
					//AuthorAvatar: DefaultAvatar
					AuthorName: admin.Name,
				}
			} else {
				// 发帖人是用户
				user, err := model.GetUserByID(info.ID)
				if err != nil {
					// 用户已注销
					m[post.AuthorID] = struct {
						AuthorName   string
						AuthorAvatar string
					}{
						AuthorAvatar: "",
						AuthorName:   "",
					}
					continue
				}
				info.AuthorName = user.Nickname
				info.AuthorAvatar = user.Avatar
				m[post.AuthorID] = struct {
					AuthorName   string
					AuthorAvatar string
				}{
					AuthorAvatar: user.Avatar,
					AuthorName:   user.Nickname,
				}
			}
		}
		postInfo[i] = info
	}
	return postInfo, nil
}
