package config

import (
	"final-project/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine{
	router := gin.Default()

	// users
	router.POST("/users/register", controllers.RegisterUser)
	router.POST("/users/login", controllers.LoginUser)
	router.PUT("/users", controllers.Tes)
	router.DELETE("/users", controllers.Tes)

	// photos
	router.POST("/photos", controllers.Tes)
	router.GET("/photos", controllers.Tes)
	router.GET("/photos/:photoId", controllers.Tes)
	router.PUT("/photos/:photoId", controllers.Tes)
	router.DELETE("/photos/:photoId", controllers.Tes)

	// comments
	router.POST("/comments", controllers.Tes)
	router.GET("/comments", controllers.Tes)
	router.GET("/comments/:commentId", controllers.Tes)
	router.PUT("/comments/:commentId", controllers.Tes)
	router.DELETE("/comments/:commentId", controllers.Tes)

	// social media
	router.POST("/socialmedias", controllers.Tes)
	router.GET("/socialmedias", controllers.Tes)
	router.PUT("/socialmedias/:socialMediaId", controllers.Tes)
	router.DELETE("/socialmedias/:socialMediaId", controllers.Tes)
	
	return router
}