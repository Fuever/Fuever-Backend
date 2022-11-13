package middleware

import (
	"Fuever/util/secret"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func failedResponse(code int, msg string) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": gin.H{},
	}
}

func Auth(ctx *gin.Context) {
	idWithTokenString := ctx.GetHeader("Authorization")
	if idWithTokenString == "" {
		ctx.AbortWithStatusJSON(200, failedResponse(40001, "without authorization"))
		return
	}
	arr := strings.Split(idWithTokenString, "@")
	if len(arr) < 2 {
		ctx.AbortWithStatusJSON(200, failedResponse(40001, "illegal token"))
		return
	}
	userID, err := strconv.Atoi(arr[0])
	if err != nil {
		ctx.AbortWithStatusJSON(200, failedResponse(40001, "illegal token"))
		return
	}
	token := arr[1]
	// key不存在或者已经过期
	if !secret.Authentication(userID, token) {
		ctx.AbortWithStatusJSON(200, failedResponse(40001, "user not in login status"))
		return
	}
	ctx.Set("userID", userID)
}
