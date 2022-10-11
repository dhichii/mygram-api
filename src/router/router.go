package router

import (
	"mygram-api/src/adapter"
	"mygram-api/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	handler := adapter.Init()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to MyGram API"})
	})

	user := router.Group("/users")
	user.POST("/register", handler.User.RegisterUserHandler)
	user.POST("/login", handler.User.LoginUserHandler)
	user.PUT("/:userId", middleware.Authentication())
	user.DELETE("/:userId", middleware.Authentication())

	photo := router.Group("/photos").Use(middleware.Authentication())
	photo.POST("", handler.Photo.PostPhotoHandler)
	photo.GET("", handler.Photo.GetAllPhotosHandler)
	photo.PUT("/:photoId", handler.Photo.UpdatePhotoHandler)
	photo.DELETE("/:photoId", handler.Photo.DeletePhotoHandler)

	return router
}
