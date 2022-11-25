package router

import (
	"Fuever/model"
	"Fuever/util/repassword"
	"math/rand"
	"strconv"
	"time"
)

// GenerateTest create demo data
func GenerateTest() {
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
		StudentID:    32004122,
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

	for i := 0; i < 17; i++ {
		model.CreateBlock(&model.Block{
			Title:    strconv.Itoa(rand.Int()),
			AuthorID: 2000000000,
		})
	}
	//TODO 说起来 评论更新帖子最后更新时间我还没写
	for i := 0; i < 114; i++ {
		model.CreatePost(&model.Post{
			AuthorID:    1,
			Title:       strconv.Itoa(rand.Int()),
			CreatedTime: time.Now().Unix(),
			UpdatedTime: time.Now().Unix(),
			State:       0,
			BlockID:     rand.Intn(17),
			IsLock:      false,
		})
	}

	for i := 0; i < 500; i++ {
		model.CreateMessage(&model.Message{
			AuthorID:    1,
			Content:     strconv.Itoa(rand.Int()),
			PostID:      rand.Intn(114),
			CreatedTime: time.Now().Unix(),
		})
	}

	for i := 0; i < 17; i++ {
		model.CreateAnniversary(&model.Anniversary{
			AdminID:   2000000000,
			Title:     strconv.Itoa(rand.Int()),
			Content:   strconv.Itoa(rand.Int()),
			Start:     rand.Int63(),
			End:       rand.Int63(),
			PositionX: rand.Float64(),
			PositionY: rand.Float64(),
		})
	}

	for i := 0; i < 27; i++ {
		model.CreateNews(&model.News{
			AuthorID:   2000000000,
			Title:      strconv.Itoa(rand.Int()),
			Content:    strconv.Itoa(rand.Int()),
			CreateTime: rand.Int63(),
			Cover:      "???",
		})
	}

}
