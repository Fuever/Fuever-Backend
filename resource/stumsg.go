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

var _studentMessageArray []*StudentMessage
var _studentMessageMap map[int]*StudentMessage
var _dormitoryMessageMap map[string][]*StudentMessage //以宿舍划分学生
var _buildingMessageMap map[int][]*StudentMessage     // 用楼号划分学生

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
	_studentMessageArray = messages
}

func initDormitoryMessageArray() {
	if _studentMessageArray == nil { //保证学生信息已初始化
		initStudentMessageArray()
	}
	_dormitoryMessageMap = make(map[string][]*StudentMessage) //初始化map
	for _, student := range _studentMessageArray {
		key := GenerateDormitoryMapKey(student.Building, student.Dormitory)
		_dormitoryMessageMap[key] = append(_dormitoryMessageMap[key], student)
	}
}

func GenerateDormitoryMapKey(building int, dormitory int) string {
	return strconv.Itoa(building) + " " + strconv.Itoa(dormitory)
}

// StudentMessages
// 维护学号到信息的映射
func StudentMessages() map[int]*StudentMessage {
	if _studentMessageMap != nil {
		return _studentMessageMap
	}
	if _studentMessageArray == nil {
		initStudentMessageArray()
	}
	_studentMessageMap = make(map[int]*StudentMessage, len(_studentMessageArray))
	for _, stu := range _studentMessageArray {
		_studentMessageMap[stu.Number] = stu
	}
	return _studentMessageMap
}

func DormitoryMessages() map[string][]*StudentMessage {
	if _dormitoryMessageMap == nil {
		initDormitoryMessageArray()
	}
	return _dormitoryMessageMap
}

// BuildingMessages
// 维护楼号到学生的映射
func BuildingMessages() map[int][]*StudentMessage {
	if _buildingMessageMap != nil {
		return _buildingMessageMap
	}
	if _studentMessageArray == nil {
		initStudentMessageArray()
	}
	_buildingMessageMap = make(map[int][]*StudentMessage, 0)
	for _, stu := range _studentMessageArray {
		_buildingMessageMap[stu.Building] = append(_buildingMessageMap[stu.Building], stu)
	}
	return _buildingMessageMap
}
