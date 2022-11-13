package service

import (
	"Fuever/util/secret"
)

// Login 返回Token
func Login(userID int) string {
	return secret.GenerateTokenAndCache(userID)
}

func Logout(userID int) {
	secret.RemoveTokenFromCache(userID)
}
