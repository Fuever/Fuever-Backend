package main

import (
	"Fuever/model"
	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDB()
	g := gin.Default()
	g.Run(":8080")
}
