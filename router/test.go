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
		Name:     "野兽先辈",
		Password: repassword.GeneratePasswordHash("114514"),
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
		Residence:    "枝江",
	})
	if err != nil {
		panic(err)
	}
	err = model.CreateUser(&model.User{
		Mail:         "123456",
		Password:     repassword.GeneratePasswordHash("12345"),
		Username:     "chen jia long",
		Nickname:     "龙除",
		Avatar:       "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com/1669256053941570786.jpg",
		StudentID:    32004106,
		Phone:        1145141919,
		Gender:       true,
		Age:          21,
		Job:          "击而破之",
		EntranceTime: time.Now().Unix(),
		Residence:    "福州",
	})
	if err != nil {
		panic(err)
	}
	err = model.CreateUser(&model.User{
		Mail:         "123457",
		Password:     repassword.GeneratePasswordHash("12345"),
		Username:     "小姨",
		Nickname:     "老姨",
		Avatar:       "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com/1669256053941570786.jpg",
		StudentID:    32004120,
		Phone:        7878787,
		Gender:       true,
		Age:          90,
		Job:          "地主",
		EntranceTime: time.Now().Unix(),
		Residence:    "地里",
	})
	if err != nil {
		panic(err)
	}
	err = model.CreateUser(&model.User{
		Mail:         "123458",
		Password:     repassword.GeneratePasswordHash("12345"),
		Username:     "鱼除",
		Nickname:     "打个胶先",
		Avatar:       "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com/1669256053941570786.jpg",
		StudentID:    32004132,
		Phone:        114514191981,
		Gender:       true,
		Age:          22,
		Job:          "胶工",
		EntranceTime: time.Now().Unix(),
		Residence:    "孙吧",
	})
	if err != nil {
		panic(err)
	}

	blockNumber := 5
	postNumber := 300
	commentNumber := 5000
	classNumber := 10
	stuNumArray := []int{32004122, 32004106, 32004132, 32004120}

	for i := 0; i < blockNumber; i++ {
		model.CreateBlock(&model.Block{
			Title:    strconv.Itoa(rand.Int()),
			AuthorID: 2000000000,
		})
	}

	for i := 0; i < postNumber; i++ {
		model.CreatePost(&model.Post{
			AuthorID:    rand.Intn(4) + 1,
			Title:       strconv.Itoa(rand.Int()),
			CreatedTime: time.Now().Unix() + int64(rand.Intn(114)*int(time.Second.Seconds())),
			UpdatedTime: time.Now().Unix() + int64(rand.Intn(114)*int(time.Second.Seconds())),
			State:       rand.Intn(3) / 2 * 2,
			BlockID:     rand.Intn(blockNumber),
			IsLock:      false,
		})
	}

	for i := 0; i < commentNumber; i++ {
		model.CreateMessage(&model.Message{
			AuthorID:    rand.Intn(4) + 1,
			Content:     strconv.Itoa(rand.Int()),
			PostID:      rand.Intn(postNumber),
			CreatedTime: time.Now().Unix() + int64(rand.Intn(114)*int(time.Second.Seconds())),
		})
	}

	for i := 0; i < 17; i++ {
		model.CreateAnniversary(&model.Anniversary{
			AdminID:   2000000000,
			Title:     "我是标题" + strconv.Itoa(rand.Int()),
			Content:   "我是标题" + strconv.Itoa(rand.Int()) + "\n" + strconv.Itoa(rand.Int()) + "\n" + strconv.Itoa(rand.Int()),
			Start:     time.Now().Unix() + int64(rand.Intn(114)*int(time.Second.Seconds())),
			End:       time.Now().Unix() + int64(rand.Intn(114)*int(time.Second.Seconds())),
			PositionX: rand.Float64(),
			PositionY: rand.Float64(),
		})
	}

	for i := 0; i < 27; i++ {
		bytes := make([]byte, 100)
		rand.Read(bytes)
		contents := make([]byte, 100)
		rand.Read(contents)
		model.CreateNews(&model.News{
			AuthorID:   2000000000,
			Title:      "我是标题" + strconv.Itoa(rand.Int()),
			Content:    "我是标题" + strconv.Itoa(rand.Int()) + "\n" + strconv.Itoa(rand.Int()) + "\n" + strconv.Itoa(rand.Int()),
			CreateTime: time.Now().Unix() + int64(rand.Intn(114)*int(time.Second.Seconds())),
			Cover:      "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com/1669256053941570786.jpg",
		})
	}

	for i := 0; i < 17; i++ {
		gallery := &model.Gallery{
			AuthorID:   2000000000,
			Title:      "我是标题" + strconv.Itoa(rand.Int()),
			Content:    "我是标题" + strconv.Itoa(rand.Int()) + "\n" + strconv.Itoa(rand.Int()) + "\n" + strconv.Itoa(rand.Int()),
			Cover:      "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com/1669256053941570786.jpg",
			CreateTime: time.Now().Unix() + int64(rand.Intn(114)*int(time.Second.Seconds())),
			PostID:     rand.Intn(100),
			PositionX:  rand.Float64(),
			PositionY:  rand.Float64(),
		}
		model.CreateGallery(gallery)
	}
	for i := 0; i < classNumber; i++ {
		clsName := strconv.Itoa(rand.Int())
		model.CreateClass(&model.Class{
			ClassName: clsName,
			StudentID: stuNumArray[rand.Intn(4)],
		})
		model.CreateClass(&model.Class{
			ClassName: clsName,
			StudentID: stuNumArray[rand.Intn(4)],
		})
		model.CreateClass(&model.Class{
			ClassName: clsName,
			StudentID: stuNumArray[rand.Intn(4)],
		})
		model.CreateClass(&model.Class{
			ClassName: clsName,
			StudentID: stuNumArray[rand.Intn(4)],
		})
	}

}
