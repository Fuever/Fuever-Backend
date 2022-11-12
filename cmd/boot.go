package main

import (
	"Fuever/model"
	"Fuever/router"
	"Fuever/util/secret"
	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDB()
	g := gin.Default()

	router.InitRoute(g)
	model.InitDB()
	secret.InitTokenCache()

	err := g.Run(":8080")
	if err != nil {
		panic(err)
	}
}
