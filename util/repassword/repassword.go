package repassword

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
)

// TODO 登录鉴权那边忘记写TODO了 先写在这

func CheckPasswordHash(password string, passwordHash string) bool {
	// 这里的passwordHash值必须要用下方的GeneratePasswordHash来生成
	// 因为我懒得判断错误哩
	arr := strings.SplitN(passwordHash, "$", 2)
	salt := arr[0]
	truePassword := arr[1]
	return truePassword == calculateHash(password, salt)
}

func GeneratePasswordHash(password string) string {
	salt := randomSalt()
	return salt + "$" + calculateHash(password, salt)
}

func calculateHash(password string, salt string) string {
	h := password
	for i := 0; i < 7; i++ {
		s := md5.Sum(append([]byte(h), []byte(salt)...))
		h = hex.EncodeToString(s[:])
	}
	return h
}

func randomSalt() string {
	bytes := make([]byte, 8)
	for i := 0; i < len(bytes); i++ {
		bytes[i] = randomString[rand.Intn(len(randomString))]
	}
	return string(bytes)
}

const randomString = "?;.|,1234567890.+-*=-!@#%^&*~_qwertyuiopasdfghjlkzmxncvb+"
