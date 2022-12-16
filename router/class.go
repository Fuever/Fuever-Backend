package router

import (
	"Fuever/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

const MaxJoin = 3

type JoinClassRequest struct {
	ClassName string `json:"class_name" binding:"required"`
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
	user, err := model.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	studentID := user.StudentID
	cnt, err := model.CountStudentJoinedClassNumber(studentID)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if cnt >= MaxJoin {
		// 加入的班级超过上限
		// 拒绝加入
		ctx.JSON(http.StatusConflict, gin.H{})
		return
	}
	err = model.CreateClass(&model.Class{
		ClassName: req.ClassName,
		StudentID: studentID,
	})
	// 这个班级已经加入过了
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

type GetClassListRequest struct {
	Offset int `form:"offset,default=0"`
	Limit  int `form:"limit" binding:"required"`
}

func GetAllClassList(ctx *gin.Context) {
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
		res[i] = &GetStudentListByClassNameResponseInfo{
			ID:        users[i].ID,
			StudentID: users[i].StudentID,
			Username:  users[i].Username,
		}
	}
	if len(res) == 0 {
		// 班级不存在
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
	return
}

func GetClassListByStudentIDWithUserAuth(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	classes, err := model.GetClassesByStudentID(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	nameArray := make([]string, len(classes))
	for i := 0; i < len(classes); i++ {
		nameArray[i] = classes[i].ClassName
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": nameArray,
	})
	return
}

type GetClassNameListByFuzzyQueryRequest struct {
	Word string `form:"word" binding:"required"`
}

func GetClassNameListByFuzzyQuery(ctx *gin.Context) {
	req := &GetClassNameListByFuzzyQueryRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	classes, err := model.GetClassesByFuzzyQuery(req.Word)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	res := make([]string, len(classes))
	for i, c := range classes {
		res[i] = c.ClassName
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
	return
}

type GetClassListByStudentIDWithAdminAuthRequest struct {
	UserID int `uri:"id" binding:"required"`
}

func GetClassListByStudentIDWithAdminAuth(ctx *gin.Context) {
	req := &GetClassListByStudentIDWithAdminAuthRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	classes, err := model.GetClassesByStudentID(req.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	res := make([]string, len(classes))
	for i, c := range classes {
		res[i] = c.ClassName
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
	return
}
