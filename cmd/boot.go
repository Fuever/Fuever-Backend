package main

import (
	"Fuever/model"
	"Fuever/router"
	"Fuever/util/secret"
	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDB()
	model.CreateAnniversary(&model.Anniversary{
		AdminID:  114514,
		Title:    "???",
		Content:  "!!!",
		Start:    1919,
		End:      810,
		Location: "???",
		Cover:    "???",
	})
	g := gin.Default()

	router.InitRoute(g)
	model.InitDB()
	secret.InitTokenCache()

	err := g.Run(":8080")
	if err != nil {
		panic(err)
	}
}
