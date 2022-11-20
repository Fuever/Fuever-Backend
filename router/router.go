package router

import (
	"Fuever/middleware"
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
				user.GET("/", nil)
				user.POST("/", nil)
				user.PUT("/", nil)
				user.DELETE("/", nil)
			}

			admin := auth.Group("/admin", nil)
			{
				//TODO
				admin.POST("/anniversary", CreateAnniversary)
				admin.PUT("/anniversary", UpdateAnniversary)
				admin.DELETE("/anniversary", DeleteAnniversary)
				admin.POST("/news", CreateNews)
				admin.PUT("/news", UpdateNews)
				admin.DELETE("/news", DeleteNews)
			}

			// recommend alumnus api
			reco := auth.Group("/recommend")
			{
				reco.GET("/", nil)
			}

			post := auth.Group("/posts")
			user.Use(middleware.UserAuth)
			{
				// return post list
				post.GET("/", nil)
				// return all message of the post which id = <:id>
				// List[Message]
				post.GET("/:id", nil)
				// create a new post
				post.POST("/", CreatePost)
				// create new message of the post which id = <:id>
				post.POST("/:id", nil)
				// delete post which id = <:id>
				post.DELETE("/:id", nil)

			}

		}

		pub := api.Group("/pub")
		{
			user := pub.Group("/user")
			{
				user.GET("/captcha", Captcha)
				user.GET("/login", UserLogin)
				user.POST("/register", Register)
				user.POST("/verify", SendEmailVerifyCode)
			}

			admin := pub.Group("/admin")
			{
				admin.POST("/login", AdminLogin)
			}

			news := pub.Group("/news")
			{
				news.GET("/", GetAllNews)
				news.GET("/:id", GetSpecifyNews)
			}

			anniv := pub.Group("/anniv")
			{
				anniv.GET("", GetAllAnniversaries)
				anniv.GET("/:id", GetSpecifyAnniversary)
			}

		}

	}

}
