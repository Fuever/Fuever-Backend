package router

import (
	"Fuever/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	BlockID int    `json:"block_id" binding:"required"`
}

func CreatePost(ctx *gin.Context) {
	req := CreatePostRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	authorID := ctx.GetInt("userID")
	post := &model.Post{
		AuthorID:    authorID,
		Title:       req.Title,
		CreatedTime: time.Now().Unix(),
		UpdatedTime: time.Now().Unix(),
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

type GetAllPostsRequest struct {
	Offset  int `uri:"offset" binding:"required"`
	Limit   int `uri:"limit" binding:"required"`
	BlockID int `uri:"block_id" binding:"required"`
}

func GetAllPosts(ctx *gin.Context) {
	req := GetAllPostsRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	posts, err := model.GetNormalPostsWithOffsetLimit(req.BlockID, req.Offset, req.Limit)
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
