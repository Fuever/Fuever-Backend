package router

import (
	"Fuever/model"
	. "Fuever/router/status"
	"Fuever/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RegisterUserInfo struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	CaptchaID       string `json:"captcha_d" binding:"required"`
	CaptchaCode     string `json:"captcha_code" binding:"required"`
	EmailVerifyCode string `json:"verify_code" binding:"required"`
}
type UserInfo struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

func ReponseWrapper(c *gin.Context, code int, msg string, data *gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// ! usage: <img src="data:image/png;base64,${imageStr}"/>
type img struct {
	Id     string `json:"captchaId" binding:"required"`
	Imgstr string `json:"imgStr" binding:"required"`
}

func GenerateAuthcode(c *gin.Context) {
	id, imgstr, _ := service.MakeCaptcha()
	ni := img{}
	ni.Id = id
	ni.Imgstr = imgstr
	// ni := gin.H{
	// 	"id":     id,
	// 	"imgstr": imgstr,
	// }
	c.JSON(http.StatusOK, gin.H{
		"code": FU_StatusOK,
		"msg":  "usage: <img src=\"data:image/png;base64,${imageStr}\"/>",
		"data": ni,
	})
}

func Register(c *gin.Context) {
	ui := RegisterUserInfo{}
	if err := c.ShouldBindJSON(&ui); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
			"data": nil,
		})
	}
	if err := ui.CaptchaCode; err != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_BadCredentials,
			"msg":  "Invalid credentials",
			"data": nil,
		})
	}
	// ! todo:add email verify function,field: RegisterUserInfo.EmailVerifyCode
	if flag := service.VerifyCaptcha(ui.CaptchaID, ui.CaptchaCode); flag == true {
		newuser := model.User{}
		newuser.Nickname = ui.Username
		newuser.Password = ui.Password
		if err := model.CreateUser(&newuser); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "Could not create the new user",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_CaptchaAuthFailed,
			"msg":  "Captcha authentication failed",
		})
	}
}

func Login(c *gin.Context) {
	ui := UserInfo{}
	if err := c.ShouldBindJSON(&ui); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  FU_ReqNotJson,
			"error": "The request body is not json-formatted",
		})
	}
	if err := ui.CaptchaCode; err != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  FU_BadCredentials,
			"error": "Invalid credentials",
		})
	}
	//  ! compare the credentials with the info in db
	//  ! hash the password & return a token or sth
	// ! restore the token to redis ?
}

func LoginAdmin(c *gin.Context) {
	ui := UserInfo{}
	if err := c.ShouldBindJSON(&ui); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	}
	if err := ui.CaptchaCode; err != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_BadCredentials,
			"msg":  "Invalid credentials",
		})
	}
	//  ! same as the above
}

func GetNews(c *gin.Context) {
	author := c.PostForm("author")
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	if author == "" {
		if newsl, err := model.GetNewsWithOffsetLimit(offset, limit); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "coule not get the news",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
				"data": newsl,
			})
		}
	} else {
		authorid, _ := strconv.Atoi(author)
		if newsl, err := model.GetNewsByAuthorIDWIthOffsetLimit(authorid, offset, limit); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "coule not get the news",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
				"data": newsl,
			})
		}
	}
}

func GetOneNews(c *gin.Context) {
	newsid, _ := strconv.Atoi(c.Query("newsid"))
	if news, err := model.GetNewByID(newsid); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "Could not get the news you requested",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
			"data": news,
		})
	}
}

func GetAnniv(c *gin.Context) {
	author := c.PostForm("author")
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	if author == "" {
		if newsl, err := model.GetAnniversariesWithOffsetLimit(offset, limit); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "Could not get the annivs info",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
				"data": newsl,
			})
		}
	} else {
		authorid, _ := strconv.Atoi(author)
		if newsl, err := model.GetNewsByAuthorIDWIthOffsetLimit(authorid, offset, limit); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "coule not get the annivs info",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
				"data": newsl,
			})
		}
	}
}

func GetOneAnniv(c *gin.Context) {
	annivid, _ := strconv.Atoi(c.Query("annivid"))
	if anniv, err := model.GetNewByID(annivid); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "Could not get the anniv you requested",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
			"data": anniv,
		})
	}
}
