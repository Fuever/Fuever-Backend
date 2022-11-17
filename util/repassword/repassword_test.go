package repassword

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func RandInt(lim int) int { //随机生成一个0~lim-1的整数
	rand.Seed(time.Now().UnixNano()) //以时间戳作为随机参数，保证结果随机
	return rand.Int() % lim
}

func RandomGenerateString(Len int) string {
	var text []byte
	for i := 0; i < Len; i++ {
		text = append(text, byte(RandInt(126)))
	}
	return string(text)
}

func TestRepassword(t *testing.T) {
	password := RandomGenerateString(10 + RandInt(10)) //随机生成10~20位的密码
	salt := RandomGenerateString(32)                   //随机生成32位的盐
	fmt.Println("password:" + password)
	fmt.Println("salt:" + salt)
	if !ProofingHashes(SaltHash(password, salt), password, salt) {
		t.Error("匹配失败")
	}
}
