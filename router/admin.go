package router

import (
	"Fuever/model"
	"Fuever/util/repassword"
	"Fuever/util/secret"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AdminLoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AdminLogin(ctx *gin.Context) {
	req := &AdminLoginRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	admin, err := model.GetAdminByName(req.Name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if !repassword.CheckPasswordHash(req.Password, admin.Password) {
		// 密码不正确
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	token := secret.GenerateTokenAndCache(admin.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"admin_id": admin.ID,
			"token":    token,
		},
	})
	return
}
