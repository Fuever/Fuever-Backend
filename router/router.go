package router

import (
	"Fuever/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(g *gin.Engine) {

	g.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	api := g.Group("/api")
	{
		auth := api.Group("/auth")
		{
			user := auth.Group("/user")
			user.Use(middleware.UserAuth)
			{
				user.GET("/:id", GetUserInfo)
				user.POST("/", nil)
				user.PUT("/", nil)
				user.DELETE("/", DeleteUser)
				user.POST("/avatar", UserUploadAvatar)
				user.DELETE("/logout", UserLogout)
			}

			admin := auth.Group("/admin")
			admin.Use(middleware.AdminAuth)
			{
				//TODO
				admin.POST("/anniversary", CreateAnniversary)
				//admin.PUT("/anniversary", UpdateAnniversary)
				admin.DELETE("/anniversary", DeleteAnniversary)
				admin.POST("/news", CreateNews)
				//admin.PUT("/news", UpdateNews)
				admin.DELETE("/news", DeleteNews)

				gallery := admin.Group("/gallery")
				{
					gallery.POST("/", CreateGallery)
					gallery.DELETE("/", DeleteGallery)
				}

				admin.POST("/block", CreateNewBlock)

				admin.POST("/img", UploadImage)
			}

			// recommend alumnus api
			reco := auth.Group("/recommend")
			{
				reco.GET("/", nil)
			}

			post := auth.Group("/posts")
			post.Use(middleware.UserAuth)
			{
				post.POST("/p/", CreatePost)
				post.POST("/p/:id", CreateComment)
				post.PUT("/p/:id", UpdateSpecifyPost)
				post.DELETE("/p/:id", DeleteSpecifyPost)
			}

		}

		pub := api.Group("/pub")
		{
			user := pub.Group("/user")
			{
				user.GET("/captcha", Captcha)
				user.POST("/login", UserLogin)
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

			gallery := pub.Group("/gallery")
			{
				gallery.GET("/", GetAllGalleries)
				gallery.GET("/:id", GetSpecifyGallery)
			}

			post := pub.Group("/posts")
			{
				post.GET("/b/:block_id", GetAllPosts)
				post.GET("/p/:id", GetSpecifyPost)
			}

			block := pub.Group("/block")
			{
				block.GET("/", GetBlockList)
			}

		}
		api.POST("test/test/test", GenerateTest)
	}

}
