package resource

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

type StudentMessage struct {
	Number    int // 特别注明一下 032004xxx会变成32004xxx 所以要保证相同的时候这里最好是用string 所以我用int
	Name      string
	Building  int
	Dormitory int
	Bed       int // A -> 1, B -> 2, C -> 3, D -> 4
}

//go:embed stu.dat
var studentMessageString string

var studentMessageArray []*StudentMessage
var dormitoryMessageArray map[string][]*StudentMessage //以宿舍划分学生

func initStudentMessageArray() {
	lines := strings.Split(studentMessageString, "\n")
	messages := make([]*StudentMessage, len(lines))
	for i := 0; i < len(lines); i++ {
		arr := strings.Split(lines[i], ",")
		studentNumber, err := strconv.Atoi(arr[0])
		if err != nil {
			log.Fatalln(err)
		}
		buildingNumber, err := strconv.Atoi(strings.Split(arr[2], "号楼")[0])
		if err != nil {
			log.Fatalln(err)
		}
		dormitoryNumber, err := strconv.Atoi(arr[3])
		if err != nil {
			log.Fatalln(err)
		}
		bedNumber := int(arr[4][0] - 'A' + 1)
		messages[i] = &StudentMessage{
			Number:    studentNumber,
			Name:      arr[1],
			Building:  buildingNumber,
			Dormitory: dormitoryNumber,
			Bed:       bedNumber,
		}
	}
	studentMessageArray = messages
}

func GenerateHash(student StudentMessage) string {
	return strconv.Itoa(student.Building) + " " + strconv.Itoa(student.Dormitory)
}

func initDormitoryMessageArray() {
	if studentMessageArray == nil { //保证学生信息已初始化
		initStudentMessageArray()
	}
	dormitoryMessageArray = make(map[string][]*StudentMessage) //初始化map
	for _, student := range studentMessageArray {
		key := GenerateHash(*student)
		dormitoryMessageArray[key] = append(dormitoryMessageArray[key], student)
	}
}

func StudentMessages() []*StudentMessage {
	if studentMessageArray == nil {
		initStudentMessageArray()
	}
	return studentMessageArray
}

func DormitoryMessages() map[string][]*StudentMessage {
	if dormitoryMessageArray == nil {
		initDormitoryMessageArray()
	}
	return dormitoryMessageArray
}
