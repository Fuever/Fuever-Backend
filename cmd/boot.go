package main

import (
	"Fuever/model"
	"Fuever/router"
	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDB()
	g := gin.Default()

	router.InitRoute(g)
	model.InitDB()

	err := g.Run(":8080")
	if err != nil {
		panic(err)
	}
}
