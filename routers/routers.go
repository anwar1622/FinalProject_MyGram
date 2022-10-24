package routers

import (
	"MyGram/controllers"
	"MyGram/middleware"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middleware.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middleware.Authentication(), controllers.DeleteUser)
	}
	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAutorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAutorization(), controllers.DeletePhoto)
	}
	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controllers.DeleteComment)
	}
	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.SocialMediaAuthorization())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}
	return router
}
