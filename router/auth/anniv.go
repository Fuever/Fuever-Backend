package auth

import (
	"Fuever/model"
	. "Fuever/router/status"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddAnnivInfo struct {
	AdminID  int    `json:"adminID"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Start    int64  `json:"start"`
	End      int64  `json:"end"`
	Location string `json:"location"`
	Cover    string `json:"cover"`
}

func AddAnniv(c *gin.Context) {
	new := AddAnnivInfo{}
	if err := c.ShouldBindJSON(&new); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		newAnniv := model.Anniversary{
			AdminID:  new.AdminID,
			Title:    new.Title,
			Content:  new.Content,
			Start:    new.Start,
			End:      new.End,
			Location: new.Location,
			Cover:    new.Cover,
		}
		if err_ := model.CreateAnniversary(&newAnniv); err_ != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "could not create the anniversary info.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	}
}
func FixAnniv(c *gin.Context) {
	alter := AddAnnivInfo{}
	if err := c.ShouldBindJSON(&alter); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		alterAnniv := model.Anniversary{
			AdminID:  alter.AdminID,
			Title:    alter.Title,
			Content:  alter.Content,
			Start:    alter.Start,
			End:      alter.End,
			Location: alter.Location,
			Cover:    alter.Cover,
		}
		if err_ := model.UpdateAnniversaryByID(&alterAnniv); err_ != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "could not update the anniversary info",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	}
}
func DelAnniv(c *gin.Context) {
	annivId, _ := strconv.Atoi(c.Query("annivId"))
	if err := model.DeleteUserByID(annivId); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "could not delete the anniversary",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
		})
	}
}
