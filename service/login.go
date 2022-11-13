package service

import (
	"Fuever/util/secret"
	"strconv"
	"strings"
)

// Login 返回Token
// 这里的token实际上是
// userID@token的形式
// 我不知道这算不算一个馊主意
func Login(userID int) string {
	return strconv.Itoa(userID) + "@" + secret.GenerateTokenAndCache(userID)
}

// Authentication
// 给定Token 判断是否登录
// 若已登录 返回userID
func Authentication(idWithTokenString string) (bool, int) {
	if idWithTokenString == "" {
		return false, 0
	}
	arr := strings.Split(idWithTokenString, "@")
	if len(arr) < 2 {
		return false, 0
	}
	userID, err := strconv.Atoi(arr[0])
	if err != nil {
		return false, 0
	}
	token := arr[1]
	// key不存在或者已经过期
	if !secret.Authentication(userID, token) {
		return false, 0
	}
	return true, userID
}

func Logout(userID int) {
	secret.RemoveTokenFromCache(userID)
}
