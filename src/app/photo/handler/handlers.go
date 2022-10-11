package handler

import (
	"mygram-api/src/app/photo"
	"mygram-api/src/app/photo/handler/request"
	"mygram-api/src/app/photo/handler/response"
	"mygram-api/src/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Handler struct {
	service photo.Service
}

func (h *Handler) PostPhotoHandler(c *gin.Context) {
	request := request.Request{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.PostPhoto(request.MapToRecord(userId))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, response.MapToPostResponse(*result))
}

func (h *Handler) GetAllPhotosHandler(c *gin.Context) {
	result, err := h.service.GetAllPhotos()
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.MapToBatchPhotoResponse(result))
}

func (h *Handler) UpdatePhotoHandler(c *gin.Context) {
	request := request.Request{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.UpdatePhoto(id, request.MapToRecord(userId))
	if err != nil {
		if err.Error() == helper.FORBIDDEN {
			helper.CreateMessageResponse(c, http.StatusForbidden,
				http.StatusText(http.StatusForbidden))
			return
		}

		if err.Error() == helper.NOTFOUND {
			helper.CreateMessageResponse(c, http.StatusNotFound,
				http.StatusText(http.StatusNotFound))
			return
		}

		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.MapToUpdateResponse(*result))
}

func (h *Handler) DeletePhotoHandler(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeletePhoto(id, userId); err != nil {
		if err.Error() == helper.FORBIDDEN {
			helper.CreateMessageResponse(c, http.StatusForbidden,
				http.StatusText(http.StatusForbidden))
			return
		}

		if err.Error() == helper.NOTFOUND {
			helper.CreateMessageResponse(c, http.StatusNotFound,
				http.StatusText(http.StatusNotFound))
			return
		}

		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	helper.CreateMessageResponse(c, http.StatusOK, "Your photo has been successfully deleted")
}

func NewHandler(service photo.Service) *Handler {
	return &Handler{service}
}
