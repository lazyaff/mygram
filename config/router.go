package config

import (
	"final-project/controllers"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine{
	router := gin.Default()

	// users
	router.POST("/users/register", controllers.RegisterUser)
	router.POST("/users/login", controllers.LoginUser)
	router.PUT("/users", middleware.Auth(), controllers.UpdateUser)
	router.DELETE("/users", middleware.Auth(), controllers.DeleteUser)

	// photos
	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Auth())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetAllUsersPhotos)
		photoRouter.GET("/:photoId", controllers.GetUsersPhoto)
		photoRouter.PUT("/:photoId", controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	}

	// comments
	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Auth())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetAllUsersComments)
		commentRouter.GET("/:commentId", controllers.GetUserComment)
		commentRouter.PUT("/:commentId", controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", controllers.DeleteComment)
	}

	// social media
	socialMediaRouter := router.Group("/socialmedias") 
	{
		socialMediaRouter.Use(middleware.Auth())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetAllSocialMedias)
		socialMediaRouter.GET("/:socialMediaId", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
	}
	
	return router
}