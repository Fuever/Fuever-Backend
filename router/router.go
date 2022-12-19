package router

import (
	"Fuever/middleware"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRoute(g *gin.Engine) {

	//GenerateTest()

	store := persistence.NewInMemoryStore(1 * time.Second)

	g.Use(middleware.Cors)
	api := g.Group("/api")
	{
		auth := api.Group("/auth")
		{
			user := auth.Group("/user")
			user.Use(middleware.UserAuth)
			{
				user.GET("/:id", GetUserInfo)
				user.GET("/r", GetUserID)
				user.PUT("/", UserUpdateInfo)
				user.DELETE("/", DeleteUser)
				user.POST("/avatar", UserUploadAvatar)
				user.DELETE("/logout", UserLogout)

				user.GET("/stu/auth", GetStudentAuthMessage)
				user.POST("/stu/auth", AuthStudentIdentity)

				user.GET("/reco/", RecommendUser)

				cls := user.Group("/cls")
				cls.Use(middleware.StuAuth)
				{
					cls.GET("/search/", GetClassNameListByFuzzyQuery)
					cls.GET("/all/", GetAllClassList)
					cls.GET("/", GetClassListByStudentIDWithUserAuth)
					cls.GET("/:name", GetStudentListByClassName)
					cls.POST("/", JoinClass)
				}
			}

			admin := auth.Group("/admin")
			admin.Use(middleware.AdminAuth)
			{

				_user := admin.Group("/user")
				{
					_user.GET("/", GetBatchUserInfo)
					_user.PUT("/", AdminUpdateUserInfo)
					_user.DELETE("/:id", AdminDeleteUser)
				}

				anniv := admin.Group("/anniv")
				{
					anniv.POST("/", CreateAnniversary)
					anniv.PUT("/", UpdateAnniversary)
					anniv.DELETE("/:id", DeleteAnniversary)
				}

				posts := admin.Group("/posts")
				{
					posts.PUT("/", ChangePostState)
					posts.DELETE("/:id", DeletePost)
				}

				comment := admin.Group("/comment")
				{
					comment.DELETE("/:id", DeleteComment)
				}

				news := admin.Group("/news")
				{
					news.GET("/", GetAllNewsWithContent)
					news.POST("/", CreateNews)
					news.PUT("/", UpdateNews)
					news.DELETE("/:id", DeleteNews)
				}

				gallery := admin.Group("/gallery")
				{
					gallery.POST("/", CreateGallery)
					gallery.DELETE("/:id", DeleteGallery)
				}

				block := admin.Group("/block")
				{
					block.POST("/", CreateNewBlock)
					block.PUT("/", UpdateBlock)
					block.DELETE("/:id", DeleteBlock)

				}

				cls := admin.Group("cls")
				{
					cls.GET("/:id", GetClassListByStudentIDWithAdminAuth)
				}

				img := admin.Group("/img")
				{
					img.POST("/", UploadImage)
				}
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
				news.GET("/", cache.CachePage(store, 1*time.Second, GetAllNews))
				news.GET("/:id", GetSpecifyNews)
			}

			anniv := pub.Group("/anniv")
			{
				anniv.GET("", cache.CachePage(store, 1*time.Second, GetAllAnniversaries))
				anniv.GET("/:id", GetSpecifyAnniversary)
			}

			gallery := pub.Group("/gallery")
			{
				gallery.GET("/", cache.CachePage(store, 1*time.Second, GetAllGalleries))
				gallery.GET("/:id", GetSpecifyGallery)
			}

			post := pub.Group("/posts")
			{
				post.GET("/search", SearchPost)
				post.GET("/", GetAllPosts)
				post.GET("/b/:block_id", GetPostsWithBlockID)
				post.GET("/p/:id", GetSpecifyPost)
			}

			block := pub.Group("/block")
			{
				block.GET("/", GetBlockList)
			}

		}
	}

}

func InitRouteWithoutCache(g *gin.Engine) {

	GenerateTest()

	g.Use(middleware.Cors)
	api := g.Group("/api")
	{
		auth := api.Group("/auth")
		{
			user := auth.Group("/user")
			user.Use(middleware.UserAuth)
			{
				user.GET("/:id", GetUserInfo)
				user.GET("/r", GetUserID)
				user.PUT("/", UserUpdateInfo)
				user.DELETE("/", DeleteUser)
				user.POST("/avatar", UserUploadAvatar)
				user.DELETE("/logout", UserLogout)

				user.GET("/stu/auth", GetStudentAuthMessage)
				user.POST("/stu/auth", AuthStudentIdentity)

				user.GET("/reco", RecommendUser)

				cls := user.Group("/cls")
				cls.Use(middleware.StuAuth)
				{
					cls.GET("/search/", GetClassNameListByFuzzyQuery)
					cls.GET("/all/", GetAllClassList)
					cls.GET("/", GetClassListByStudentIDWithUserAuth)
					cls.GET("/:name", GetStudentListByClassName)
					cls.POST("/", JoinClass)
				}
			}

			admin := auth.Group("/admin")
			admin.Use(middleware.AdminAuth)
			{

				_user := admin.Group("/user")
				{
					_user.GET("/", GetBatchUserInfo)
					_user.PUT("/", AdminUpdateUserInfo)
					_user.DELETE("/:id", AdminDeleteUser)
				}

				anniv := admin.Group("/anniv")
				{
					anniv.POST("/", CreateAnniversary)
					anniv.PUT("/", UpdateAnniversary)
					anniv.DELETE("/:id", DeleteAnniversary)
				}

				posts := admin.Group("/posts")
				{
					posts.PUT("/", ChangePostState)
					posts.DELETE("/:id", DeletePost)
				}

				comment := admin.Group("/comment")
				{
					comment.DELETE("/:id", DeleteComment)
				}

				news := admin.Group("/news")
				{
					news.POST("/", CreateNews)
					news.PUT("/", UpdateNews)
					news.DELETE("/:id", DeleteNews)
				}

				gallery := admin.Group("/gallery")
				{
					gallery.POST("/", CreateGallery)
					gallery.DELETE("/:id", DeleteGallery)
				}

				block := admin.Group("/block")
				{
					block.POST("/", CreateNewBlock)
					block.PUT("/", UpdateBlock)
					block.DELETE("/:id", DeleteBlock)

				}

				cls := admin.Group("cls")
				{
					cls.GET("/:id", GetClassListByStudentIDWithAdminAuth)
				}

				img := admin.Group("/img")
				{
					img.POST("/", UploadImage)
				}
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
				post.GET("/search", SearchPost)
				post.GET("/", GetAllPosts)
				post.GET("/b/:block_id", GetPostsWithBlockID)
				post.GET("/p/:id", GetSpecifyPost)
			}

			block := pub.Group("/block")
			{
				block.GET("/", GetBlockList)
			}

		}
	}

}
