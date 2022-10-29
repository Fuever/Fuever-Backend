package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInfo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	AuthCode string `json:"authcode" binding:"required"`
}

func GenerateAuthcode(c *gin.Context) {
	// ! TODO
}

func Register(c *gin.Context) {
	ri := RegisterInfo{}
	if err := c.ShouldBindJSON(&ri); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"The request body is not json-formatted",
		})
	}
	if err :=ri.AuthCode;err !="" {
		c.JSON(http.StatusForbidden,gin.H{
			"error":"Invalid credentials",
		})
	}
	// ! TODO
}