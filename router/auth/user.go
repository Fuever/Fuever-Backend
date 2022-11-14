package auth

import (
	"Fuever/model"
	. "Fuever/router/status"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type alterUserInfo struct {
	Mail         string `json:"mail"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avator"`
	StudentID    int    `json:"student_id"`
	Phone        int    `json:"phone"`
	Gender       bool   `json:"gender"`
	Age          int    `json:"age"`
	Job          string `json:"job"`
	EntranceTime int64  `json:"entrance_time"`
	ClassID      int    `json:"class_id"`
	Residence    string `json:"residence"`
}

func GetUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.DefaultQuery("userId", ""))
	if user, err := model.GetUserByID(userId); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "Could not get the user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
			"data": user,
		})
	}
}

func UpdateUser(c *gin.Context) {
	alter := alterUserInfo{}
	if err := c.ShouldBindJSON(&alter); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		alterUser := model.User{
			Mail:         alter.Mail,
			Username:     alter.Username,
			Nickname:     alter.Nickname,
			Avatar:       alter.Avatar,
			StudentID:    alter.StudentID,
			Phone:        alter.Phone,
			Gender:       alter.Gender,
			Age:          alter.Age,
			Job:          alter.Job,
			EntranceTime: alter.EntranceTime,
			ClassID:      alter.ClassID,
			Residence:    alter.Residence,
		}
		if err_ := model.UpdateUser(&alterUser); err_ != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "could not update the user info",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	}
}
func AddUser(c *gin.Context) {

}

// func AddUser(c *gin.Context) {
// 	new := model.User{}
// 	if err := c.ShouldBindJSON(&new); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": FU_ReqNotJson,
// 			"msg":  "The request body is not json-formatted",
// 		})
// 	} else {
// 		if err_ := model.CreateUser(&new); err_ != nil {
// 			c.JSON(http.StatusOK, gin.H{
// 				"code": FU_DBError,
// 				"msg":  "could not create the user",
// 			})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{
// 				"code": FU_StatusOK,
// 				"msg":  "ok",
// 			})
// 		}
// 	}
// }

func DelUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	if err := model.DeleteUserByID(userId); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "could not delete the user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
		})
	}
}
