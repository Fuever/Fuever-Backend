package router

import (
	"Fuever/model"
	"Fuever/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetAllGalleries(ctx *gin.Context) {
	galleries, err := service.GetAllGalleries()
	if err != nil && err != gorm.ErrRecordNotFound {
		// 不是因为not found没法找到记录而报错
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": galleries,
	})
	return
}

type GetSpecifyGalleryRequest struct {
	ID int `uri:"id" binding:"required"`
}

func GetSpecifyGallery(ctx *gin.Context) {
	req := &GetSpecifyGalleryRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	galleryID := req.ID
	info, err := service.GetSpecifyGallery(galleryID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": info,
	})
	return
}

type CreateGalleryRequest struct {
	Title     string  `json:"title" binding:"required"`
	Content   string  `json:"content" binding:"required"`
	Cover     string  `json:"cover" binding:"required"`
	PositionX float64 `json:"position_x" binding:"required"`
	PositionY float64 `json:"position_y" binding:"required"`
}

func CreateGallery(ctx *gin.Context) {
	authID := ctx.GetInt("authID")
	req := &CreateGalleryRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	post := &model.Post{
		AuthorID:    authID,
		Title:       req.Title,
		CreatedTime: time.Now().Unix(),
		UpdatedTime: time.Now().Unix(),
		State:       0,
		BlockID:     0,
		IsLock:      false,
	}
	err := model.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	gallery := &model.Gallery{
		AuthorID:   authID,
		Title:      req.Title,
		Content:    req.Content,
		Cover:      req.Cover,
		CreateTime: time.Now().Unix(),
		PostID:     post.ID,
		PositionX:  req.PositionX,
		PositionY:  req.PositionY,
	}
	err = model.CreateGallery(gallery)
	if err != nil {
		// 如果title已存在 会报错 回滚创建
		model.DeletePostByID(gallery.PostID)
		ctx.JSON(http.StatusConflict, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type DeleteGalleryRequest struct {
	ID int `json:"id" binding:"required"`
}

func DeleteGallery(ctx *gin.Context) {
	adminID := ctx.GetInt("adminID")
	req := &DeleteGalleryRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	gallery, err := model.GetGalleryByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	if gallery.AuthorID != adminID {
		// 不是创建者 不允许删除
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	err = model.DeleteGalleryID(req.ID)
	if err != nil {
		// 找不到记录
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}
