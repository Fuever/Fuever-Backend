package router

import (
	"Fuever/model"
	"Fuever/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type GetBlockListRequest struct {
	Offset int `form:"offset,default=0"`
	Limit  int `form:"limit" binding:"required"`
}

func GetBlockList(ctx *gin.Context) {
	req := &GetBlockListRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	blocks, err := service.GetBlocks(req.Limit, req.Offset)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": blocks,
	})
	return
}

type CreateNewBlockRequest struct {
	Title string `json:"title" binding:"required"`
}

// CreateNewBlock 只有管理员可以创建主题(板块)
func CreateNewBlock(ctx *gin.Context) {
	adminID := ctx.GetInt("adminID")
	req := &CreateNewBlockRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err := model.CreateBlock(&model.Block{
		Title:    req.Title,
		AuthorID: adminID,
	})
	if err != nil {
		// 如果这条记录已经存在
		ctx.JSON(http.StatusConflict, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type UpdateBlockRequest struct {
	ID    int    `json:"id"    binding:"required"`
	Title string `json:"title" binding:"required"`
}

func UpdateBlock(ctx *gin.Context) {
	req := &UpdateBlockRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	block, err := model.GetBlockByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	block.Title = req.Title
	if err := model.UpdateBlock(block); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type DeleteBlockRequest struct {
	ID int `json:"id"    binding:"required"`
}

func DeleteBlock(ctx *gin.Context) {
	req := &DeleteBlockRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	block, err := model.GetBlockByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	err = model.DeleteBlockByID(block.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}
