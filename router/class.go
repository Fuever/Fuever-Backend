package router

//TODO 测试这三个接口

import (
	"Fuever/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

const MaxJoin = 3

type JoinClassRequest struct {
	ClassName string `form:"class_name" binding:"required"`
}

// JoinClass
//
//	创建和加入就不区分了
func JoinClass(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	req := &JoinClassRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	cnt, err := model.CountStudentJoinedClassNumber(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if cnt > MaxJoin {
		// 加入的班级超过上限
		// 拒绝加入
		ctx.JSON(http.StatusConflict, gin.H{})
		return
	}
	err = model.CreateClass(&model.Class{
		ClassName: req.ClassName,
		StudentID: userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type GetClassListRequest struct {
	Offset int `form:"offset,limit=0" binding:"required"`
	Limit  int `form:"limit" binding:"required"`
}

func GetClassList(ctx *gin.Context) {
	req := &GetClassListRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	arr, err := model.GetClassList(req.Offset, req.Limit)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	classArray := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		classArray[i] = arr[i].ClassName
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": classArray,
	})
	return
}

type GetStudentListByClassNameRequest struct {
	ClassName string `uri:"name" binding:"required"`
}

type GetStudentListByClassNameResponseInfo struct {
	ID        int    `json:"id"`
	StudentID int    `json:"student_id"`
	Username  string `json:"name"`
}

func GetStudentListByClassName(ctx *gin.Context) {
	req := &GetStudentListByClassNameRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	className := req.ClassName
	users, err := model.GetStudentListByClassName(className)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	res := make([]*GetStudentListByClassNameResponseInfo, len(users))
	for i := 0; i < len(users); i++ {
		res[i].ID = users[i].ID
		res[i].StudentID = users[i].StudentID
		res[i].Username = users[i].Username
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
	return
}
