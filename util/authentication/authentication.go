package authentication

import (
	"Fuever/resource"
	"math/rand"
)

// GenerateStudentAuthMessage
// 根据学号返回验证信息
// @return  待验证的一组姓名
func GenerateStudentAuthMessage(studentNumber int, studentName string) ([]string, bool) {
	stuNameMap := make(map[string]bool, 0)

	stuMessage, flag := resource.StudentMessages()[studentNumber]
	if !flag {
		// 学号不存在于这个表中
		return nil, false
	}
	if stuMessage.Name != studentName {
		// 名字和学号对不上
		return nil, false
	}
	key := resource.GenerateDormitoryMapKey(stuMessage.Building, stuMessage.Dormitory)
	roommateArray, flag := resource.DormitoryMessages()[key]
	if !flag {
		// 这人是没舍友的家伙啊 我能怎么办呢
		// TODO 这里可能要开一个管理员审核通道
		return nil, false
	}
	for _, roommate := range roommateArray {
		// 注意不要直接用原来的数组
		// 会改变资源里的值
		// 说到底到底为什么不能immutable啊
		if roommate.Name == stuMessage.Name {
			// 本人不在验证集合的考虑范围内
			continue
		}
		stuNameMap[roommate.Name] = true
	}
	stuArray, flag := resource.BuildingMessages()[stuMessage.Building]
	if !flag || len(stuArray) < 13 {
		// 这栋楼是什么啊？
		// 不住人吗
		return nil, false
	}
	for len(stuNameMap) < 16 {
		randomStu := stuArray[rand.Int()%len(stuArray)]
		stuNameMap[randomStu.Name] = true
	}
	res := make([]string, 0)
	for k := range stuNameMap {
		res = append(res, k)
	}
	return res, true
}

func CheckStudentAuthMessage(studentNumber int, studentName string, roommates []string) bool {
	stuMessage, flag := resource.StudentMessages()[studentNumber]
	if !flag {
		// 学号不存在于这个表中
		return false
	}
	if stuMessage.Name != studentName {
		// 名字和学号对不上
		return false
	}
	key := resource.GenerateDormitoryMapKey(stuMessage.Building, stuMessage.Dormitory)
	correctRoommatesArray := resource.DormitoryMessages()[key]
	if len(correctRoommatesArray)-1 != len(roommates) {
		// 除去本人外人数都对不上
		return false
	}
	correctRoommatesMap := make(map[string]bool, 0)
	for _, roommate := range correctRoommatesArray {
		correctRoommatesMap[roommate.Name] = true
	}
	roommatesRoommatesMap := make(map[string]bool, 0)
	for _, roommate := range roommates {
		roommatesRoommatesMap[roommate] = true
	}
	if len(correctRoommatesMap)-1 != len(roommatesRoommatesMap) {
		// 避免重名
		return false
	}
	for k := range roommatesRoommatesMap {
		if _, exist := correctRoommatesMap[k]; !exist {
			return false
		}
	}
	return true
}
