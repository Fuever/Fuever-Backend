package authentication

import (
	"Fuever/resource"
	"fmt"
	"strconv"
	"testing"
)

func RandStudent(array []*resource.StudentMessage) resource.StudentMessage {
	return *array[RandInt(len(array))] //随机返回一位学生的信息
}

func DisplayStudentMessage(student resource.StudentMessage) {
	number := strconv.Itoa(student.Number)       //学号
	name := student.Name                         //姓名
	build := strconv.Itoa(student.Building)      //楼号
	dormitory := strconv.Itoa(student.Dormitory) //宿舍号
	bed := strconv.Itoa(student.Bed)             //床号
	fmt.Println(number + "," + name + "," + build + "," + dormitory + "," + bed)
}

func TestAutentication(t *testing.T) {
	//dormitoryMessageArray := resource.DormitoryMessages()
	//for _, array := range dormitoryMessageArray {
	//	for i := 0; i < len(array); i++ {
	//		DisplayStudentMessage(*array[i])
	//	}
	//	fmt.Println("----------------")
	//}
	//for i := 0; i < 50; i++ {
	//	rand.Seed(time.Now().UnixNano())
	//	x := rand.Int()
	//	fmt.Println(x)
	//}
	studentMessageArray := resource.StudentMessages() //获取学生信息
	vistor := RandStudent(studentMessageArray)        //随机生成学生
	fmt.Println("待验证访客")
	DisplayStudentMessage(vistor) //输出访客信息
	//生成随机学生列表并输出
	fmt.Println("验证名单...")
	randomList := GenerateAuthenticationList(vistor)
	for _, student := range randomList {
		DisplayStudentMessage(*student) //输出随机名单
	}
}
