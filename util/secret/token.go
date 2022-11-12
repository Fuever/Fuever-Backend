package secret

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

var (
	redisClient *redis.Client
	ctx         context.Context
)

const (
	secretKey  = "veni vidi vici"
	expireTime = 37 * time.Minute
)

// GenerateTokenAndCache
// 根据输入的用户ID生成Token
// 并且缓存到redis中
// 以 <userID : token > 的形式存储
func GenerateTokenAndCache(userID int) string {
	token := generateToken(userID)
	err := redisClient.Set(ctx, strconv.Itoa(userID), token, expireTime).Err()
	if err != nil {
		log.Println(err)
		return ""
	}
	return token
}

// Authentication
// 根据用户的ID和传入的token验证该用户是否已经登录
// 如果已登录返回true
// 负责返回false
func Authentication(userID int, token string) bool {
	userIDString := strconv.Itoa(userID)
	realToken, err := redisClient.Get(ctx, userIDString).Result()
	if err != nil {
		log.Println(err)
		return false
	}
	// 比较token是否正确
	res := realToken == token
	if res {
		// 如果token是正确的
		// 就续期
		err = redisClient.Expire(ctx, userIDString, expireTime).Err()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	return false
}

// generateToken returns token
// 这个函数不是纯函数
func generateToken(userID int) string {
	s1 := strconv.Itoa(int(time.Now().Unix()))
	s2 := string(rune(userID))
	s := s1 + s2
	for i := 0; i < 3; i++ {
		m := md5.New()
		m.Write([]byte(s + secretKey))
		s = hex.EncodeToString(m.Sum(nil))
	}
	return s
}

// InitTokenCache 初始化redis连接
func InitTokenCache() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx = context.Background()
}
