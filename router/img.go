package router

import (
	"Fuever/util/img"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
	}
	imageName := img.SaveImage(file)
	ctx.JSON(http.StatusOK, gin.H{
		"data": imageName,
	})
}
