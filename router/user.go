package router

import (
	"Fuever/model"
	"Fuever/service"
	"Fuever/util/img"
	"Fuever/util/repassword"
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
	ctx.JSON(http.StatusOK, gin.H{})
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

type UserLoginRequest struct {
	Mailbox  string `json:"mailbox" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserLogin(ctx *gin.Context) {
	req := &UserLoginRequest{}
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
	token := service.Login(user.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token": token,
		},
	})
	return
}

func UserLogout(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	service.Logout(userID)
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

func UserUploadAvatar(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	fileHeader, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	avatarPath := img.SaveImage(file)
	user, err := model.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	user.Avatar = avatarPath
	err = model.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"avatar": avatarPath,
	})
	return
}

type GetUserInfoRequest struct {
	ID int `uri:"id" binding:"required"`
}

type UserInfo struct {
	ID           int    `json:"id"`
	Mail         string `json:"mail"`
	Nickname     string `json:"nickname,omitempty"`
	Username     string `json:"username,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	StudentID    int    `json:"student_id,omitempty"`
	Phone        int    `json:"phone,omitempty"`
	Gender       bool   `json:"gender,omitempty"`
	Age          int    `json:"age,omitempty"`
	Job          string `json:"job,omitempty"`
	EntranceTime int64  `json:"entrance_time,omitempty"`
	ClassID      int    `json:"class_id,omitempty"`
	Residence    string `json:"residence,omitempty"`
}

func GetUserInfo(ctx *gin.Context) {
	req := &GetUserInfoRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	user, err := model.GetUserByID(req.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	info := &UserInfo{
		ID:           user.ID,
		Mail:         user.Mail,
		Nickname:     user.Nickname,
		Username:     user.Username,
		Avatar:       user.Avatar,
		StudentID:    user.StudentID,
		Phone:        user.Phone,
		Gender:       user.Gender,
		Age:          user.Age,
		Job:          user.Job,
		EntranceTime: user.EntranceTime,
		ClassID:      user.ClassID,
		Residence:    user.Residence,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": info,
	})
	return
}
