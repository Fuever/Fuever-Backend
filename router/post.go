package router

import (
	"Fuever/model"
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

type GetAllPostsUriRequest struct {
	BlockID int `uri:"block_id" binding:"required"`
}

type GetAllPostsQueryRequest struct {
	Offset int `form:"offset" binding:"required"`
	Limit  int `form:"limit" binding:"required"`
}

func GetAllPosts(ctx *gin.Context) {
	uriReq := &GetAllPostsUriRequest{}
	if err := ctx.ShouldBindUri(uriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	urlReq := &GetAllPostsQueryRequest{}
	if err := ctx.ShouldBindQuery(urlReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	posts, err := model.GetNormalPostsWithOffsetLimit(uriReq.BlockID, urlReq.Offset, urlReq.Limit)
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

func GetSpecifyPost(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, gin.H{
		"data": post,
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
