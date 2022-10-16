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
	user.PUT("/:userId", middleware.Authentication(), handler.User.UpdateUserHandler)
	user.DELETE("/:userId", middleware.Authentication(), handler.User.DeleteUserHandler)

	photo := router.Group("/photos").Use(middleware.Authentication())
	photo.POST("", handler.Photo.PostPhotoHandler)
	photo.GET("", handler.Photo.GetAllPhotosHandler)
	photo.PUT("/:photoId", handler.Photo.UpdatePhotoHandler)
	photo.DELETE("/:photoId", handler.Photo.DeletePhotoHandler)

	socialMedia := router.Group("/socialmedias").Use(middleware.Authentication())
	socialMedia.POST("", handler.SocialMedia.CreateSocialMediaHandler)
	socialMedia.GET("", handler.SocialMedia.GetAllSocialMediasHandler)
	socialMedia.PUT("/:socialMediaId", handler.SocialMedia.UpdateSocialMediaHandler)
	socialMedia.DELETE("/:socialMediaId", handler.SocialMedia.DeleteSocialMediaHandler)

	comment := router.Group("/comments").Use(middleware.Authentication())
	comment.POST("", handler.Comment.CreateCommentHandler)
	comment.GET("", handler.Comment.GetAllCommentsHandler)
	comment.PUT("/:commentId", handler.Comment.UpdateCommentHandler)
	comment.DELETE("/:commentId", handler.Comment.DeleteCommentHandler)

	return router
}
