package router

import (
	"Fuever/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.DefaultQuery("userid", ""))
	if user, err := model.GetUserByID(userid); err != nil {
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
	newone := model.User{}
	if err := c.ShouldBindJSON(&newone); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		if err_ := model.UpdateUser(&newone); err_ != nil {
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
	ui := model.User{}
	if err := c.ShouldBindJSON(&ui); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		if err_ := model.CreateUser(&ui); err_ != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "could not create the user",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	}
}

func DelUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Query("userid"))
	if err := model.DeleteUserByID(userid); err != nil {
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
func AddAnniv(c *gin.Context) {
	na := model.Anniversary{}
	if err := c.ShouldBindJSON(&na); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		if err_ := model.CreateAnniversary(&na); err_ != nil {
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
	newone := model.Anniversary{}
	if err := c.ShouldBindJSON(&newone); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		if err_ := model.UpdateAnniversaryByID(&newone); err_ != nil {
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
	annivid, _ := strconv.Atoi(c.Query("annivid"))
	if err := model.DeleteUserByID(annivid); err != nil {
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
func AddNews(c *gin.Context) {
	nn := model.New{}
	if err := c.ShouldBindJSON(&nn); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		if err_ := model.CreateNew(&nn); err_ != nil {
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
	newone := model.New{}
	if err := c.ShouldBindJSON(&newone); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		if err_ := model.UpdateNew(&newone); err_ != nil {
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
	newsid, _ := strconv.Atoi(c.Query("annivid"))
	if err := model.DeleteNewByID(newsid); err != nil {
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

func GetPosts(c *gin.Context) {

}
func GetReviews(c *gin.Context) {

}
func CreatePost(c *gin.Context) {

}
func ReviewPost(c *gin.Context) {

}
func DeletePost(c *gin.Context) {

}
