package service

import (
	"Fuever/util/mail"
	"fmt"
	"math/rand"
	"time"
)

const checkTime = 70 * time.Second
const expireTime = 3 * time.Minute

type verifyCodeInfo struct {
	mailbox string
	expired int64
}

// verifyCode 映射到verifyCodeInfo
var verifyCodeMap = make(map[int]*verifyCodeInfo, 0)

// mailbox 映射到 上一封发给这个邮箱的邮件的时间
var hasMailboxBeenSent = make(map[string]int64, 0)

var content = "your verify code is %v"

func VerifyCode(mailbox string, verifyCode int) bool {
	info, flag := verifyCodeMap[verifyCode]
	if !flag {
		// 如果不存在这个信息
		// 说明码错了
		return false
	}
	if info.expired < time.Now().Unix() {
		//已经过期
		return false
	}
	if info.mailbox != mailbox {
		// 如果邮箱和对应验证码的不同
		// 说明🐴还是错的
		return false
	}
	delete(verifyCodeMap, verifyCode)
	delete(hasMailboxBeenSent, mailbox)
	return true
}

func SendVerifyCodeToUserMailbox(mailbox string) error {
	if t, flag := hasMailboxBeenSent[mailbox]; flag && t > time.Now().Unix() {
		// 当之前发送过且🐴还没过期时进入该选择肢
		// 之前发送过了
		// 不能频繁的发送
		// 因为没有米
		return nil
	}
	verifyCode := generateMailboxVerifyCode()
	hasMailboxBeenSent[mailbox] = time.Now().Unix() + expireTime.Milliseconds()
	verifyCodeMap[verifyCode] = &verifyCodeInfo{
		mailbox: mailbox,
		expired: time.Now().Unix() + expireTime.Milliseconds(),
	}
	return mail.SendEmail(mailbox, getContent(verifyCode))
}

func getContent(param ...any) string {
	return fmt.Sprintf(content, param)
}

func generateMailboxVerifyCode() int {
	for true {
		code := rand.Int()
		if code/100000 == 0 {
			// 不是六位有效数字
			continue
		}
		if _, flag := verifyCodeMap[code]; flag {
			// 如果这个verify code已经存在
			// 需要重新生成
			continue
		} else {
			return code
		}

	}
	return -1
}

func init() {
	go func() {
		ticker := time.NewTicker(checkTime)
		// 删除过期键值对
		for range ticker.C {
			for k, v := range verifyCodeMap {
				if v.expired < time.Now().Unix() {
					delete(verifyCodeMap, k)
				}
			}
			for k, v := range hasMailboxBeenSent {
				if v < time.Now().Unix() {
					delete(hasMailboxBeenSent, k)
				}
			}
		}
	}()
}
