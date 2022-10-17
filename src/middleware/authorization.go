package middleware

import (
	"mygram-api/src/app/auth"
	"mygram-api/src/app/auth/repository"
	"mygram-api/src/app/auth/service"
	"mygram-api/src/database"
	"mygram-api/src/helper"
	"mygram-api/src/helper/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type middleware struct {
	service auth.Service
}

func newAuthmiddleware(service auth.Service) *middleware {
	return &middleware{service}
}

func InitAuthmiddleware() *middleware {
	db := database.GetPostgresConnection()
	repo := repository.NewGORMRepository(db)
	serv := service.NewService(repo)
	return newAuthmiddleware(serv)
}

func (m *middleware) PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
			return
		}

		userId, err := m.service.GetUserIdByPhotoId(photoId)
		if err != nil {
			if err.Error() == helper.NOTFOUND {
				c.AbortWithStatusJSON(
					http.StatusNotFound,
					structs.Message{
						Message: http.StatusText(http.StatusNotFound)},
				)
				return
			}

			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				structs.Message{
					Message: http.StatusText(http.StatusInternalServerError)},
			)
			return
		}

		userData := helper.GetUserData(c)
		if userId != userData.ID {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				http.StatusText(http.StatusForbidden),
			)
		}

		c.Next()
	}
}

func (m *middleware) CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
			return
		}

		userId, err := m.service.GetUserIdByCommentId(commentId)
		if err != nil {
			if err.Error() == helper.NOTFOUND {
				c.AbortWithStatusJSON(
					http.StatusNotFound,
					structs.Message{
						Message: http.StatusText(http.StatusNotFound)},
				)
				return
			}

			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				structs.Message{
					Message: http.StatusText(http.StatusInternalServerError)},
			)
			return
		}

		userData := helper.GetUserData(c)
		if userId != userData.ID {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				http.StatusText(http.StatusForbidden),
			)
		}

		c.Next()
	}
}

func (m *middleware) SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, structs.Message{Message: err.Error()})
			return
		}

		userId, err := m.service.GetUserIdBySocialMediaId(socialMediaId)
		if err != nil {
			if err.Error() == helper.NOTFOUND {
				c.AbortWithStatusJSON(
					http.StatusNotFound,
					structs.Message{
						Message: http.StatusText(http.StatusNotFound)},
				)
				return
			}

			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				structs.Message{
					Message: http.StatusText(http.StatusInternalServerError)},
			)
			return
		}

		userData := helper.GetUserData(c)
		if userId != userData.ID {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				http.StatusText(http.StatusForbidden),
			)
		}

		c.Next()
	}
}
