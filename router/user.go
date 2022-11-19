package router

import (
	"Fuever/model"
	"Fuever/service"
	"Fuever/util/repassword"
	"Fuever/util/secret"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type EmailVerifyCodeRequest struct {
	VerifyID   string `json:"verify_id" binding:"required"`
	VerifyCode string `json:"verify_code" binding:"required"`
	Mailbox    string `json:"mailbox" binding:"required"`
}

func SendEmailVerifyCode(ctx *gin.Context) {
	req := &EmailVerifyCodeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if !service.VerifyCaptcha(req.VerifyID, req.VerifyCode) {
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	_, err := model.GetUserByMailbox(req.Mailbox)
	if err != gorm.ErrRecordNotFound {
		// 如果这条记录可见 说明已存在该用户
		ctx.JSON(http.StatusConflict, gin.H{"msg": "user exist"})
		return
	}
	err = service.SendVerifyCodeToUserMailbox(req.Mailbox)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type RegisterRequest struct {
	Mailbox        string `json:"mailbox" binding:"required"`
	Password       string `json:"password" binding:"required"`
	MailVerifyCode int    `json:"mail_verify_code" binding:"required"`
}

func Register(ctx *gin.Context) {
	req := &RegisterRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if !service.VerifyCode(req.Mailbox, req.MailVerifyCode) {
		// 邮箱验证码不正确
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	user := &model.User{
		Mail:     req.Mailbox,
		Password: repassword.GeneratePasswordHash(req.Password),
	}
	_, err := model.GetUserByMailbox(req.Mailbox)
	if err == nil {
		// 用户已存在
		ctx.JSON(http.StatusConflict, gin.H{})
		return
	}
	err = model.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// 密码不能传到网络上啊喂
	user.Password = ""
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
	if !repassword.CheckPasswordHash(req.Password, user.Password) {
		// 密码错误
		ctx.JSON(http.StatusForbidden, gin.H{})
		return
	}
	token := secret.GenerateTokenAndCache(user.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user_id": user.ID,
			"token":   token,
		},
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
