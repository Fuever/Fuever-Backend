package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudentAuthMessage(t *testing.T) {
	for i := 0; i < MaxRetryTimes; i++ {
		res, flag := CheckStudentAuthMessage(1114514, 1919810, "野兽先辈", []string{"孙笑川"})
		assert.Equal(t, false, flag)
		assert.Equal(t, false, res)
	}
	res, flag := CheckStudentAuthMessage(1114514, 1919810, "野兽先辈", []string{"孙笑川"})
	assert.Equal(t, true, flag)
	assert.Equal(t, false, res)
	arr, flag := GenerateStudentAuthMessage(114514, 1919810, "野兽先辈")
	assert.ObjectsAreEqualValues(nil, arr)
	assert.Equal(t, false, flag)
}
