package router

import (
	. "Fuever/router/auth"

	"github.com/gin-gonic/gin"
)

func InitRoute(g *gin.Engine) {

	loginCheck := func(ctx *gin.Context) {} // authentication middleware

	api := g.Group("/api")
	{
		auth := api.Group("/auth", loginCheck)
		{
			user := auth.Group("/user", nil)
			{
				user.GET("/", GetUser)
				user.POST("/", UpdateUser)
				user.PUT("/", AddUser)
				user.DELETE("/", DelUser)
			}

			admin := auth.Group("/admin", nil)
			{
				//TODO
				admin.POST("/anniversary", AddAnniv)
				admin.PUT("/anniversary", FixAnniv)
				admin.DELETE("/anniversary", DelAnniv)
				admin.POST("/news", AddNews)
				admin.PUT("/news", FixNews)
				admin.DELETE("/news", DelNews)
			}

			// recommend alumnus api
			reco := auth.Group("/recommend")
			{
				reco.GET("/", nil)
			}

			post := auth.Group("/posts")
			{
				// return post list
				post.GET("/", GetPosts)
				// return all message of the post which id = <:id>
				// List[Message]
				post.GET("/:id", GetReviews)
				// create a new post
				post.POST("/", CreatePost)
				// create new message of the post which id = <:id>
				post.POST("/:id", ReplyPost)
				// delete post which id = <:id>
				post.DELETE("/:id", DeletePost)

			}

		}

		pub := api.Group("/pub")
		{
			user := pub.Group("/user")
			{
				user.GET("/captcha", GenerateAuthcode)
				user.POST("/login", Login)
				user.POST("/register", Register)
				user.POST("/verify",VerifyEmail)
			}

			admin := api.Group("/admin")
			{
				admin.POST("/login", LoginAdmin)
			}

			news := api.Group("/news")
			{
				news.GET("/", GetOneNews)
				news.POST("/", GetNews)
			}

			anniv := api.Group("/anniversary")
			{
				anniv.GET("/", GetOneAnniv)
				anniv.POST("/", GetAnniv)
			}

		}

	}

}
