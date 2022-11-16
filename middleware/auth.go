package middleware

import (
	"Fuever/service"
	"github.com/gin-gonic/gin"
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
		ctx.AbortWithStatusJSON(200, failedResponse(40001, "user not in login status"))
		return
	}
}
