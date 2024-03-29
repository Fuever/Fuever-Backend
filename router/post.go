package router

import (
	"Fuever/model"
	"Fuever/service"
	"Fuever/util/sensitive"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	BlockID int    `json:"block_id" binding:"required"`
}

func CreatePost(ctx *gin.Context) {
	req := &CreatePostRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	authorID := ctx.GetInt("userID")
	nowTimeUnix := time.Now().Unix()
	post := &model.Post{
		AuthorID:    authorID,
		Title:       req.Title,
		CreatedTime: nowTimeUnix,
		UpdatedTime: nowTimeUnix,
		State:       0,
		BlockID:     req.BlockID,
		IsLock:      false,
	}
	err := model.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"post_id": post.ID,
		},
	})
	return
}

type GetAllPostsQueryRequest struct {
	Offset int `form:"offset,default=0"`
	Limit  int `form:"limit" binding:"required"`
}

func GetAllPosts(ctx *gin.Context) {
	req := &GetAllPostsQueryRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	posts, err := service.GetAllPost(req.Offset, req.Limit)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
	return
}

type GetPostsWithBlockIDUriRequest struct {
	BlockID int `uri:"block_id" binding:"required"`
}

type GetPostsWithBlockIDQueryRequest struct {
	Offset int `form:"offset,default=0"`
	Limit  int `form:"limit" binding:"required"`
}

func GetPostsWithBlockID(ctx *gin.Context) {
	uriReq := &GetPostsWithBlockIDUriRequest{}
	if err := ctx.ShouldBindUri(uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	urlReq := &GetPostsWithBlockIDQueryRequest{}
	if err := ctx.ShouldBindQuery(urlReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	posts, err := service.GetPosts(uriReq.BlockID, urlReq.Offset, urlReq.Limit)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
	return
}

type SpecifyPostRequest struct {
	ID int `uri:"id" binding:"required"`
}

type GetSpecifyPostQueryRequest struct {
	Offset int `form:"offset,default=0"`
	Limit  int `form:"limit" binding:"required"`
}

func GetSpecifyPost(ctx *gin.Context) {
	uriReq := &SpecifyPostRequest{}
	if err := ctx.ShouldBindUri(uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	queryReq := &GetSpecifyPostQueryRequest{}
	if err := ctx.ShouldBindQuery(queryReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	postInfo, err := service.GetPost(uriReq.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	comments, err := service.GetComments(postInfo.ID, queryReq.Offset, queryReq.Limit)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"post":    postInfo,
			"comment": comments,
		},
	})
	return
}

type UpdateSpecifyPostRequest struct {
	NewTitle string `json:"new_title" binding:"required"`
}

// UpdateSpecifyPost 楼主仅允许修改标题
func UpdateSpecifyPost(ctx *gin.Context) {
	req := &SpecifyPostRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	updateReq := &UpdateSpecifyPostRequest{}
	if err := ctx.ShouldBindJSON(updateReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	newTitle := updateReq.NewTitle
	post, err := model.GetPostByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	userID := ctx.GetInt("userID")
	//不是作者没有修改权限
	if userID != post.AuthorID {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	//被锁定了无法修改
	if post.IsLock {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	post.Title = newTitle
	post.UpdatedTime = time.Now().Unix()
	err = model.UpdatePost(post)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

func DeleteSpecifyPost(ctx *gin.Context) {
	req := &SpecifyPostRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	post, err := model.GetPostByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	userID := ctx.GetInt("userID")
	//不是作者没有修改权限
	if userID != post.AuthorID {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	//被锁定了无法修改
	if post.IsLock {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	err = model.DeletePostByID(post.ID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

func CreateComment(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	uriReq := &SpecifyPostRequest{}
	if err := ctx.ShouldBindUri(uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	req := &CreateCommentRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	now := time.Now().Unix()
	postID := uriReq.ID
	content := sensitive.GetFilter().ReplaceSensitiveWord(req.Content, "*")
	comment := &model.Message{
		AuthorID:    userID,
		Content:     content,
		PostID:      postID,
		CreatedTime: now,
	}
	///* 这地方要查2次表 */
	//// 确保被插入帖子的评论存在
	post, err := model.GetPostByID(comment.PostID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	post.UpdatedTime = now
	err = model.UpdatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	///****************/
	//err := model.UpdatePostUpdatedTimeByID(comment.PostID, now)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{})
	//	return
	//}
	err = model.CreateMessage(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type ChangePostStateRequest struct {
	ID    int `json:"id" binding:"required"`
	State int `json:"state" binding:"required"`
}

func ChangePostState(ctx *gin.Context) {
	req := &ChangePostStateRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if req.State < 0 || req.State > 2 {
		// 不为正常状态
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	post, err := model.GetPostByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	post.State = req.State
	err = model.UpdatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type DeletePostRequest struct {
	ID int `uri:"id" binding:"required"`
}

func DeletePost(ctx *gin.Context) {
	req := &DeletePostRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	post, err := model.GetPostByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	err = model.DeletePostByID(post.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type SearchPostRequest struct {
	Word   string `form:"word" binding:"required"`
	Offset int    `form:"offset, default=0"`
	Limit  int    `form:"limit" binding:"required"`
}

func SearchPost(ctx *gin.Context) {
	req := &SearchPostRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	info, err := service.GetPostsWithFuzzyStringOffsetLimit(req.Word, req.Offset, req.Limit)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": info,
	})
	return
}

type DeleteCommentRequest struct {
	ID int `uri:"id" binding:"required"`
}

func DeleteComment(ctx *gin.Context) {
	req := &DeleteCommentRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	msg, err := model.GetMessageByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	err = model.DeleteMessageByID(msg.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}
