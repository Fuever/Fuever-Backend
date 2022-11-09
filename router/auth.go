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
	blockid, _ := strconv.Atoi(c.DefaultQuery("blockid", "0"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "10"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if posts, err := model.GetNormalPostsWithOffsetLimit(blockid, offset, limit); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "coule not get the posts",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
			"data": posts,
		})
	}
}
func GetReviews(c *gin.Context) {
	postid, _ := strconv.Atoi(c.Param("id"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "10"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if posts, err := model.GetMessageByPostIDWithOffsetLimit(postid, offset, limit); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "coule not get the post's review",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
			"data": posts,
		})
	}
}
func CreatePost(c *gin.Context) {
	np := model.Post{}
	if err := c.ShouldBindJSON(&np); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		if err_ := model.CreatePost(&np); err_ != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "could not create the post",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	}
}
func ReplyPost(c *gin.Context) {
	nm := model.Message{}
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		if err_ := model.CreateMessage(&nm); err_ != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_DBError,
				"msg":  "could not create the reply",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": FU_StatusOK,
				"msg":  "ok",
			})
		}
	}
}
func DeletePost(c *gin.Context) {
	postid, _ := strconv.Atoi(c.Query("postid"))
	if err := model.DeletePostByID(postid); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_DBError,
			"msg":  "could not delete the post",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_StatusOK,
			"msg":  "ok",
		})
	}
}
