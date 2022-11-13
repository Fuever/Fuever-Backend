package service

import (
	"Fuever/util/secret"
	"strconv"
)

// Login 返回Token
// 这里的token实际上是
// userID@token的形式
// 我不知道这算不算一个馊主意
func Login(userID int) string {
	return strconv.Itoa(userID) + "@" + secret.GenerateTokenAndCache(userID)
}

func Logout(userID int) {
	secret.RemoveTokenFromCache(userID)
}
