package auth

import (
	"Fuever/model"
	. "Fuever/router/status"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddNewInfo struct {
	AuthorID int    `json:"adminID"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Cover    string `json:"cover"`
}

func AddNews(c *gin.Context) {
	new := AddNewInfo{}
	if err := c.ShouldBindJSON(&new); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		newNew := model.New{
			AuthorID: new.AuthorID,
			Title:    new.Title,
			Content:  new.Content,
			Cover:    new.Cover,
		}
		if err_ := model.CreateNew(&newNew); err_ != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "could not create the news",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	}
}
func FixNews(c *gin.Context) {
	alter := AddNewInfo{}
	if err := c.ShouldBindJSON(&alter); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		alterNew := model.New{
			AuthorID: alter.AuthorID,
			Title:    alter.Title,
			Content:  alter.Content,
			Cover:    alter.Cover,
		}
		if err_ := model.UpdateNew(&alterNew); err_ != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "could not update the news",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	}
}
func DelNews(c *gin.Context) {
	newsId, _ := strconv.Atoi(c.Query("newsId"))
	if err := model.DeleteNewByID(newsId); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "could not delete the news",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
		})
	}
}
