package router

import (
	"Fuever/model"
	"Fuever/util/repassword"
	"github.com/gin-gonic/gin"
	"time"
)

func GenerateTest(ctx *gin.Context) {
	err := model.CreateAdmin(&model.Admin{
		Name:     "1234567",
		Password: repassword.GeneratePasswordHash("1234567"),
	})
	if err != nil {
		panic(err)
	}
	err = model.CreateUser(&model.User{
		Mail:         "12345",
		Password:     repassword.GeneratePasswordHash("12345"),
		Username:     "王日斤",
		Nickname:     "向晚大魔王",
		Avatar:       "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com/1669256053941570786.jpg",
		StudentID:    032004122,
		Phone:        11451419198,
		Gender:       true,
		Age:          22,
		Job:          "胶工",
		EntranceTime: time.Now().Unix(),
		ClassID:      114514,
		Residence:    "枝江",
	})
	if err != nil {
		panic(err)
	}
}
