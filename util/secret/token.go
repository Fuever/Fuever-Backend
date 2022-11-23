package secret

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	redisClient *redis.Client
	ctx         context.Context
	_addr       = os.Getenv("FUEVER_CACHE")
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
		// userID 对应的Token不存在
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
		m := sha512.New()
		m.Write([]byte(s + secretKey))
		s = hex.EncodeToString(m.Sum(nil))
	}
	return s
}

func RemoveTokenFromCache(userID int) {
	redisClient.Del(ctx, strconv.Itoa(userID))
}

// InitTokenCache 初始化redis连接
func InitTokenCache() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     _addr + ":6379",
		Password: "",
		DB:       0,
	})
	ctx = context.Background()
}
