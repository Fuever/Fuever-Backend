package router

import (
	"Fuever/model"
	"github.com/gin-gonic/gin"
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
