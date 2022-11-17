package router

import (
	"Fuever/model"
	"Fuever/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SendEmailVerifyCode() {

}

type RegisterRequest struct {
	VerifyID   string `json:"verify_id" binding:"required"`
	VerifyCode string `json:"verify_code" binding:"required"`
	Mailbox    string `json:"mailbox" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	req := &RegisterRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if !service.VerifyCaptcha(req.VerifyID, req.VerifyCode) {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	user := &model.User{
		Mail:     req.Mailbox,
		Password: req.Password, //TODO 加密
	}
	// TODO 邮箱验证
	err := model.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"msg": "user exist",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
	return
}

type LoginRequest struct {
	Mailbox  string `json:"mailbox" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	req := &LoginRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	user, err := model.GetUserByMailbox(req.Mailbox)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
	return
}

func Captcha(ctx *gin.Context) {
	id, imgBase64, err := service.MakeCaptcha()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"verify_id": id,
			"img":       imgBase64,
		},
	})
	return
}
