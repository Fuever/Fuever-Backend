package router

import "github.com/gin-gonic/gin"

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
				admin.POST("/anniversary", nil)
				admin.PUT("/anniversary", nil)
				admin.DELETE("/anniversary", nil)
				admin.POST("/news", nil)
				admin.PUT("/news", nil)
				admin.DELETE("/news", nil)
			}

			// recommend alumnus api
			reco := auth.Group("/recommend")
			{
				reco.GET("/", nil)
			}

			post := auth.Group("/posts")
			{
				// return post list
				post.GET("/", nil)
				// return all message of the post which id = <:id>
				// List[Message]
				post.GET("/:id", nil)
				// create a new post
				post.POST("/", nil)
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
				user.POST("/authcode", GenerateAuthcode)
				user.POST("/login", Login)
				user.POST("/register", Register)

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
