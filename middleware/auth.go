package middleware

import (
	"Fuever/model"
	"Fuever/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func failedResponse(code int, msg string) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": gin.H{},
	}
}

// UserAuth 针对需要用户登录接口
func UserAuth(ctx *gin.Context) {
	idWithTokenString := ctx.GetHeader("Authorization")
	if isLogin, userID := service.Authentication(idWithTokenString); isLogin {
		// 如果鉴权成功
		// 向后文提供userID
		// 注意 如果过了验证就信任用户传来的ID
		// 会出现token和实际请求的userID不一致的情况
		// userID请从ctx中获取 而不是在请求体中获取
		// 最好是能来一发assert(userID, requestUserID)
		ctx.Set("userID", userID)
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		return
	}
}

// AdminAuth 针对需要管理员的登录接口
func AdminAuth(ctx *gin.Context) {
	idWithTokenString := ctx.GetHeader("Authorization")
	if isLogin, adminID := service.Authentication(idWithTokenString); isLogin {
		if adminID < 2000000000 {
			// 虽然查得到TOKEN
			// 但是是普通用户
			// 拒绝请求
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		}
		ctx.Set("adminID", adminID)
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		return
	}
}

// StuAuth 针对需要过了学生验证才能访问的接口
func StuAuth(ctx *gin.Context) {
	idWithTokenString := ctx.GetHeader("Authorization")
	if isLogin, userID := service.Authentication(idWithTokenString); isLogin {
		user, err := model.GetUserByID(userID)
		if err != nil {
			// 没有这条记录
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
			return
		}
		if user.StudentID != 0 && user.Username != "" {
			// 过了学生验证
			ctx.Set("userID", userID)
			return
		} else {
			// 没有过学生验证
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
			return
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
		return
	}
}
