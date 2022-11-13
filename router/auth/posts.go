package auth

import (
	"Fuever/model"
	. "Fuever/router/status"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JsonPost struct {
	AuthorID int    `json:"authorID"`
	Title    string `json:"title"`
	BlockID  int    `json:"blockID"`
}
type JsonReview struct {
	AuthorID int    `json:"authorID"`
	Content  string `json:"content"`
	PostID   int    `json:"postID"`
}

func GetPosts(c *gin.Context) {
	blockId, _ := strconv.Atoi(c.DefaultQuery("blockId", "0"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "10"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if posts, err := model.GetNormalPostsWithOffsetLimit(blockId, offset, limit); err != nil {
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
	postId, _ := strconv.Atoi(c.Param("id"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "10"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if posts, err := model.GetMessageByPostIDWithOffsetLimit(postId, offset, limit); err != nil {
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
	new := JsonPost{}
	if err := c.ShouldBindJSON(&new); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		newPost := model.Post{
			AuthorID: new.AuthorID,
			Title:    new.Title,
			BlockID:  new.BlockID,
		}
		if err_ := model.CreatePost(&newPost); err_ != nil {
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
	reply := JsonReview{}
	if err := c.ShouldBindJSON(&reply); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": FU_ReqNotJson,
			"msg":  "The request body is not json-formatted",
		})
	} else {
		newReply := model.Message{
			AuthorID: reply.AuthorID,
			Content:  reply.Content,
			PostID:   reply.PostID,
		}
		if err_ := model.CreateMessage(&newReply); err_ != nil {
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
	postId, _ := strconv.Atoi(c.Query("postId"))
	if err := model.DeletePostByID(postId); err != nil {
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
