package router

import (
	"Fuever/model"
	"Fuever/util/repassword"
	"time"
)

// GenerateTest create demo data
func GenerateTest() {
	err := model.CreateAdmin(&model.Admin{
		Name:     "哦金金",
		Password: repassword.GeneratePasswordHash("wx114514"),
	})
	if err != nil {
		panic(err)
	}
	err = model.CreateAdmin(&model.Admin{
		Name:     "龙爷",
		Password: repassword.GeneratePasswordHash("cjl12345"),
	})
	if err != nil {
		panic(err)
	}
	err = model.CreateAdmin(&model.Admin{
		Name:     "老姨",
		Password: repassword.GeneratePasswordHash("hxy12345"),
	})
	if err != nil {
		panic(err)
	}
	err = model.CreateAdmin(&model.Admin{
		Name:     "野兽先辈",
		Password: repassword.GeneratePasswordHash("__1919810"),
	})
	if err != nil {
		panic(err)
	}
	err = model.CreateUser(&model.User{
		Mail:         "114514@1919810.com",
		Password:     repassword.GeneratePasswordHash("114514@1919810.com"),
		Username:     "王日斤",
		Nickname:     "向晚大魔王",
		Avatar:       "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com/1669256053941570786.jpg",
		StudentID:    32004122,
		Phone:        11451419198,
		Gender:       true,
		Age:          22,
		Job:          "胶工",
		EntranceTime: time.Now().Unix(),
		Residence:    "枝江",
	})
}
