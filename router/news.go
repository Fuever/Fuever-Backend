package router

import (
	"Fuever/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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
	news, err := model.GetNewsByID(req.ID)
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
	newses, err := model.GetNewsWithOffsetLimit(offset, limit)
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

// CreateNews need Admin Auth
func CreateNews(ctx *gin.Context) {

}

// UpdateNews need Admin Auth
func UpdateNews(ctx *gin.Context) {

}

// DeleteNews need Admin Auth
func DeleteNews(ctx *gin.Context) {

}
