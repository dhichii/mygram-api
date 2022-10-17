package handler

import (
	"errors"
	"mygram-api/src/app/photo"
	"mygram-api/src/app/photo/handler/request"
	"mygram-api/src/app/photo/handler/response"
	"mygram-api/src/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type Handler struct {
	service photo.Service
}

// @Tags Photo
// @Summary Post photo
// @ID post-photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body request.Request true "json request body"
// @Success 201 {object} response.PostResponse
// @Router /photos [post]
func (h *Handler) PostPhotoHandler(c *gin.Context) {
	request := request.Request{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))

	if err := c.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, helper.ValidateRequest(verr))
			return
		}

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

// @Tags Photo
// @Summary Get all photos
// @ID get-all-photos
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {array} response.PhotoResponse
// @Router /photos [get]
func (h *Handler) GetAllPhotosHandler(c *gin.Context) {
	result, err := h.service.GetAllPhotos()
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.MapToBatchPhotoResponse(result))
}

// @Tags Photo
// @Summary Update photo
// @ID update-photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "photoId"
// @Param RequestBody body request.Request true "json request body"
// @Success 200 {object} response.UpdateResponse
// @Router /photos/{photoId} [put]
func (h *Handler) UpdatePhotoHandler(c *gin.Context) {
	request := request.Request{}
	id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, helper.ValidateRequest(verr))
			return
		}

		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userData := helper.GetUserData(c)
	result, err := h.service.UpdatePhoto(id, request.MapToRecord(userData.ID))
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

// @Tags Photo
// @Summary Delete photo
// @ID delete-photo
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path int true "photoId"
// @Success 200 {object} structs.Message
// @Router /photos/{photoId} [delete]
func (h *Handler) DeletePhotoHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userData := helper.GetUserData(c)
	if err := h.service.DeletePhoto(id, userData.ID); err != nil {
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
