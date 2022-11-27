package router

import (
	"Fuever/model"
	"Fuever/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type GetNewsRequest struct {
	ID int `uri:"id" binding:"required"`
}

// GetSpecifyNews 无需认证
func GetSpecifyNews(ctx *gin.Context) {
	req := &GetNewsRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	news, err := service.GetNews(req.ID)
	if err != nil {
		// 返回为空
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": news})
	return
}

// GetAllNews 不需要任何认证
func GetAllNews(ctx *gin.Context) {
	offsetString, offsetFlag := ctx.GetQuery("offset")
	limitString, limitFlag := ctx.GetQuery("limit")
	offset, offsetErr := strconv.Atoi(offsetString)
	limit, limitErr := strconv.Atoi(limitString)
	if !offsetFlag || !limitFlag || offsetErr != nil || limitErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	newses, err := service.GetNewses(offset, limit)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果错误是记录不存在
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		// 服务器错误
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": newses})
	return
}

type CreateNewsRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Cover   string `json:"cover" binding:"required"`
}

// CreateNews need Admin Auth
func CreateNews(ctx *gin.Context) {
	adminID := ctx.GetInt("adminID")
	req := &CreateNewsRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	news := &model.News{
		AuthorID:   adminID,
		Title:      req.Title,
		Content:    req.Content,
		CreateTime: time.Now().Unix(),
		Cover:      req.Cover,
	}
	err := model.CreateNews(news)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type UpdateNewsRequest struct {
	ID      int    `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Cover   string `json:"cover" binding:"required"`
}

// UpdateNews need Admin Auth
func UpdateNews(ctx *gin.Context) {
	req := &UpdateNewsRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	news, err := model.GetNewsByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	news.Title = req.Title
	news.Content = req.Content
	news.Cover = req.Cover
	err = model.UpdateNews(news)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type DeleteNewsRequest struct {
	ID int `uri:"id" binding:"required"`
}

// DeleteNews need Admin Auth
func DeleteNews(ctx *gin.Context) {
	adminID := ctx.GetInt("adminID")
	req := &DeleteGalleryRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	news, err := model.GetNewsByID(req.ID)
	if err != nil {
		// 该条记录不存在
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	if adminID != news.AuthorID {
		// 不是作者捏
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	err = model.DeleteNewsByID(req.ID)
	if err != nil {
		// 记录不存在的话不会到这个地方捏
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}
