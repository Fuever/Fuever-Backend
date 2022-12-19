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
	//err := model.CreateAdmin(&model.Admin{
	//	Name:     "哦金金",
	//	Password: repassword.GeneratePasswordHash("wx114514"),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//err = model.CreateAdmin(&model.Admin{
	//	Name:     "龙爷",
	//	Password: repassword.GeneratePasswordHash("cjl12345"),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//err = model.CreateAdmin(&model.Admin{
	//	Name:     "老姨",
	//	Password: repassword.GeneratePasswordHash("hxy12345"),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//err = model.CreateAdmin(&model.Admin{
	//	Name:     "野兽先辈",
	//	Password: repassword.GeneratePasswordHash("__1919810"),
	//})
	//if err != nil {
	//	panic(err)
	//}
	for i := 0; i < 17; i++ {
		model.CreateUser(&model.User{
			Mail:         strconv.Itoa(i) + "14514@1919810.com",
			Password:     repassword.GeneratePasswordHash("114514@1919810.com"),
			Username:     "王日斤" + strconv.Itoa(i),
			Nickname:     "向晚大魔王",
			Avatar:       "https://fuever-1313037799.cos.ap-nanjing.myqcloud.com/1669256053941570786.jpg",
			StudentID:    32004100 - i,
			Phone:        11451419198,
			Gender:       true,
			Age:          22,
			Job:          "胶工",
			EntranceTime: time.Now().Unix(),
			Residence:    "枝江",
		})
	}

	for i := 10; i <= 22; i++ {
		model.CreateClass(&model.Class{
			ClassName: strconv.Itoa(i) + "届大数据01班",
			StudentID: 32004100 - rand.Intn(17),
		})
	}
	for i := 10; i <= 22; i++ {
		model.CreateClass(&model.Class{
			ClassName: strconv.Itoa(i) + "届计算机01班",
			StudentID: 32004100 - rand.Intn(17),
		})
		model.CreateClass(&model.Class{
			ClassName: strconv.Itoa(i) + "届计算机02班",
			StudentID: 32004100 - rand.Intn(17),
		})
		model.CreateClass(&model.Class{
			ClassName: strconv.Itoa(i) + "届计算机03班",
			StudentID: 32004100 - rand.Intn(17),
		})
		model.CreateClass(&model.Class{
			ClassName: strconv.Itoa(i) + "届计算机04班",
			StudentID: 32004100 - rand.Intn(17),
		})
	}
	for i := 10; i <= 22; i++ {
		model.CreateClass(&model.Class{
			ClassName: strconv.Itoa(i) + "届信息安全01班",
			StudentID: 32004100 - rand.Intn(17),
		})
	}
	for i := 10; i <= 22; i++ {
		model.CreateClass(&model.Class{
			ClassName: strconv.Itoa(i) + "届软件工程01班",
			StudentID: 32004100 - rand.Intn(17),
		})
	}
	for i := 17; i <= 22; i++ {
		model.CreateClass(&model.Class{
			ClassName: strconv.Itoa(i+114514-20) + "届抽象工程00班",
			StudentID: 32004100 - rand.Intn(17),
		})
	}
}
