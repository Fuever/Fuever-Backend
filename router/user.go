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
	if !service.IsLegalMailbox(req.Mailbox) {
		ctx.JSON(http.StatusBadRequest, gin.H{})
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
		ctx.JSON(http.StatusForbidden, gin.H{})
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
		Residence:    user.Residence,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": info,
	})
	return
}

type UserUpdateInfoRequest struct {
	Nickname     string `json:"nickname"`
	Phone        int    `json:"phone"`
	Gender       bool   `json:"gender"`
	Age          int    `json:"age"`
	Job          string `json:"job"`
	EntranceTime int64  `json:"entrance_time"`
	Residence    string `json:"residence"`
}

func UserUpdateInfo(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	req := &UserUpdateInfoRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	user, err := model.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	user.Nickname = req.Nickname
	user.Phone = req.Phone
	user.Gender = req.Gender
	user.Age = req.Age
	user.Job = req.Job
	user.EntranceTime = req.EntranceTime
	user.Residence = req.Residence
	err = model.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

// DeleteUser 用户自己注销
func DeleteUser(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	// 缓存要先清除
	service.Logout(userID)
	err := model.DeleteUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type GetBatchUserInfoRequest struct {
	Offset int `form:"offset, default=0"`
	Limit  int `form:"limit" binding:"required"`
}

func GetBatchUserInfo(ctx *gin.Context) {
	req := &GetBatchUserInfoRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	users, err := model.GetUserWithOffsetLimit(req.Offset, req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	usersInfo := make([]*UserInfo, len(users))
	for i := 0; i < len(users); i++ {
		usersInfo[i] = &UserInfo{
			ID:           users[i].ID,
			Mail:         users[i].Mail,
			Nickname:     users[i].Nickname,
			Username:     users[i].Username,
			Avatar:       users[i].Avatar,
			StudentID:    users[i].StudentID,
			Phone:        users[i].Phone,
			Gender:       users[i].Gender,
			Age:          users[i].Age,
			Job:          users[i].Job,
			EntranceTime: users[i].EntranceTime,
			Residence:    users[i].Residence,
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": usersInfo,
	})
	return
}

func AdminUpdateUserInfo(ctx *gin.Context) {
	req := &UserInfo{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	user, err := model.GetUserByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	user.Nickname = req.Nickname
	user.Username = req.Username
	user.Gender = req.Gender
	user.StudentID = req.StudentID
	user.Phone = req.Phone
	user.Mail = req.Mail
	user.Age = req.Age
	user.Job = req.Job
	user.Avatar = req.Avatar
	user.Residence = req.Residence
	user.EntranceTime = req.EntranceTime
	err = model.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type AdminDeleteUserRequest struct {
	ID int `uri:"id" binding:"required"`
}

func AdminDeleteUser(ctx *gin.Context) {
	req := &AdminDeleteUserRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	user, err := model.GetUserByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	// 缓存要先清除
	service.Logout(user.ID)
	err = model.DeleteUserByID(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

func GetUserID(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	ctx.JSON(http.StatusOK, gin.H{
		"data": userID,
	})
	return
}

type GetStudentAuthMessageRequest struct {
	Name          string `form:"name" binding:"required"`
	StudentNumber int    `form:"student_number" binding:"required"`
}

// GetStudentAuthMessage
// 获取验证信息
func GetStudentAuthMessage(ctx *gin.Context) {
	userID := ctx.GetInt("userID")

	req := &GetStudentAuthMessageRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	arr, flag := service.GenerateStudentAuthMessage(userID, req.StudentNumber, req.Name)
	if !flag {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"name_array": arr,
		},
	})
}

type AuthStudentIdentityRequest struct {
	Name          string   `json:"name" binding:"required"`
	StudentNumber int      `json:"student_number" binding:"required"`
	Roommates     []string `json:"roommates" binding:"required"`
}

// AuthStudentIdentity
// 验证是不是本校的学生
func AuthStudentIdentity(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	req := &AuthStudentIdentityRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	isAuthSucceed, isOverStep := service.CheckStudentAuthMessage(userID, req.StudentNumber, req.Name, req.Roommates)
	if isOverStep || !isAuthSucceed {
		// 验证失败 或是 请求太多次
		// 所以茶壶不能烧咖啡
		// 以后别再这么做了
		ctx.JSON(http.StatusTeapot, gin.H{})
		return
	}
	user, err := model.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	user.Username = req.Name
	user.StudentID = req.StudentNumber
	err = model.UpdateUser(user)
	if err != nil {
		// 这个已经被人验证过了啊?
		// 怎么又来一遍还能对啊
		// 爬
		ctx.JSON(http.StatusConflict, gin.H{})
		return
	}
	// 终于验证成功了
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

const _recommendUserNumber = 3

func RecommendUser(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	user, err := model.PickRandomUser(userID, _recommendUserNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	info := make([]*UserInfo, len(user))
	for i := 0; i < len(user); i++ {
		info[i] = &UserInfo{
			ID:        user[i].ID,
			Username:  user[i].Username,
			Nickname:  user[i].Nickname,
			Avatar:    user[i].Avatar,
			Residence: user[i].Residence,
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": info,
	})
	return
}
