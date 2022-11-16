package authentication

import (
	"Fuever/resource"
	"math/rand"
	"time"
)

func RandInt(lim int) int { //随机生成一个0~lim-1的整数
	rand.Seed(time.Now().Unix()) //以时间戳作为随机参数，保证结果随机
	return rand.Int() % lim
}

func Count(array []*resource.StudentMessage, student resource.StudentMessage) bool { //查找该学生是否已在列表内
	for _, value := range array {
		if *value == student {
			return true
		}
	}
	return false
}

func GenerateAuthenticationList(vistor resource.StudentMessage) []*resource.StudentMessage {
	studentMessageArray := resource.StudentMessages()     //获取学生信息
	dormitoryMessageArray := resource.DormitoryMessages() //获取宿舍信息
	var randomList []*resource.StudentMessage             //随机学生序列，共9位
	//处理同宿舍随机
	key := resource.GenerateHash(vistor)
	roomieList := dormitoryMessageArray[key]
	numberOfRoomie := len(roomieList)      //宿舍人数
	countInside := RandInt(numberOfRoomie) //随机同一宿舍内的舍友数量：0～3
	for i, j := 0, 0; i < numberOfRoomie && j < countInside; i++ {
		if *roomieList[i] != vistor { //避免加入验证者自身
			randomList = append(randomList, roomieList[i])
			j++ //成功加入随机列表
		}
	}
	//处理不同宿舍随机
	countOutside := 9 - countInside             //不同宿舍的学生数量,满足countInside+countOutside=9
	numberOfStudent := len(studentMessageArray) //学生数量
	for i := 0; i < countOutside; {
		index := RandInt(numberOfStudent)
		if *studentMessageArray[index] != vistor && !Count(randomList, *studentMessageArray[index]) { //避免重复加入信息
			randomList = append(randomList, studentMessageArray[index])
			i++ //成功加入随机列表
		}
	}
	//打乱随机序列
	rand.Seed(time.Now().Unix())                   //以时间戳作为随机参数，保证结果随机
	rand.Shuffle(len(randomList), func(i, j int) { //随机打乱
		randomList[i], randomList[j] = randomList[j], randomList[i]
	})
	return randomList
}
