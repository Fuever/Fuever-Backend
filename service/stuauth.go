package service

import "Fuever/util/authentication"

const MaxRetryTimes = 3

// 验证每错误一次  这里就加一条记录
// 错误次数超过MaxRetryTimes时 禁止再次验证
var _failMap = make(map[int]int, 0)

// GenerateStudentAuthMessage
// @return
// index 0: 待验证的姓名集合
// index 1: 输入的信息是否合法(学号和姓名是否匹配)
func GenerateStudentAuthMessage(userID int, studentNumber int, studentName string) ([]string, bool) {
	num, flag := _failMap[userID]
	if flag && num >= MaxRetryTimes {
		// 超过最大重试次数
		return nil, false
	}
	arr, flag := authentication.GenerateStudentAuthMessage(studentNumber, studentName)
	if !flag {
		// 学号姓名不匹配或者压根不存在
		return nil, false
	}
	return arr, true
}

// CheckStudentAuthMessage
// @return
// index 0: 验证是否成功
// index 1: 是否超过最大重试次数
func CheckStudentAuthMessage(userID int, studentNumber int, studentName string, roommates []string) (bool, bool) {
	num, flag := _failMap[userID]
	if flag && num >= MaxRetryTimes {
		// 超过最大重试次数
		return false, true
	}
	if !authentication.CheckStudentAuthMessage(studentNumber, studentName, roommates) {
		// 验证失败 记录次数
		_failMap[userID]++
		return false, false
	}
	// 验证成功 清除错误记录
	delete(_failMap, userID)
	return true, false
}
