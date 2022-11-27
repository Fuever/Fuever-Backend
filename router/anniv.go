package router

import (
	"Fuever/model"
	"Fuever/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type GetAnniversaryRequest struct {
	ID int `uri:"id" binding:"required"`
}

// GetSpecifyAnniversary 无需认证
func GetSpecifyAnniversary(ctx *gin.Context) {
	req := &GetAnniversaryRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	news, err := service.GetAnniversary(req.ID)
	if err != nil {
		// 返回为空
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": news})
	return

}

// GetAllAnniversaries 不需要任何认证
func GetAllAnniversaries(ctx *gin.Context) {
	offsetString, offsetFlag := ctx.GetQuery("offset")
	limitString, limitFlag := ctx.GetQuery("limit")
	offset, offsetErr := strconv.Atoi(offsetString)
	limit, limitErr := strconv.Atoi(limitString)
	if !offsetFlag || !limitFlag || offsetErr != nil || limitErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	anniv, err := service.GetAnniversaries(offset, limit)
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
	ctx.JSON(http.StatusOK, gin.H{"data": anniv})
	return
}

type CreateAnniversaryRequest struct {
	Title     string  `json:"title" binding:"required"`
	Content   string  `json:"content" binding:"required"`
	Start     int64   `json:"start" binding:"required"`
	End       int64   `json:"end" binding:"required"`
	PositionX float64 `json:"position_x" binding:"required"`
	PositionY float64 `json:"position_y" binding:"required"`
}

// CreateAnniversary need Admin Auth
func CreateAnniversary(ctx *gin.Context) {
	adminID := ctx.GetInt("adminID")
	req := &CreateAnniversaryRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	anniv := &model.Anniversary{
		AdminID:   adminID,
		Title:     req.Title,
		Content:   req.Content,
		Start:     req.Start,
		End:       req.End,
		PositionX: req.PositionX,
		PositionY: req.PositionY,
	}
	err := model.CreateAnniversary(anniv)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type UpdateAnniversaryRequest struct {
	ID        int     `json:"id" binding:"required"`
	Title     string  `json:"title" binding:"required"`
	Content   string  `json:"content" binding:"required"`
	Start     int64   `json:"start" binding:"required"`
	End       int64   `json:"end" binding:"required"`
	PositionX float64 `json:"position_x" binding:"required"`
	PositionY float64 `json:"position_y" binding:"required"`
}

// UpdateAnniversary need Admin Auth
func UpdateAnniversary(ctx *gin.Context) {
	req := &model.Anniversary{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	anniv, err := service.GetAnniversary(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	anniv.Title = req.Title
	anniv.Content = req.Content
	anniv.Start = req.Start
	anniv.End = req.End
	anniv.PositionX = req.PositionX
	anniv.PositionY = req.PositionY
	err = model.UpdateAnniversaryByID(&anniv.Anniversary)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type DeleteAnniversaryRequest struct {
	ID int `uri:"id" binding:"required"`
}

// DeleteAnniversary need Admin Auth
func DeleteAnniversary(ctx *gin.Context) {
	adminID := ctx.GetInt("adminID")
	req := &DeleteGalleryRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	anniv, err := model.GetAnniversaryByID(req.ID)
	if err != nil {
		// 该条记录不存在
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	if adminID != anniv.AdminID {
		// 不是作者捏
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	err = model.DeleteAnniversaryByID(req.ID)
	if err != nil {
		// 记录不存在的话不会到这个地方捏
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}
