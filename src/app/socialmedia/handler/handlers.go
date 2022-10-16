package handler

import (
	"errors"
	"mygram-api/src/app/socialmedia"
	"mygram-api/src/app/socialmedia/handler/request"
	"mygram-api/src/app/socialmedia/handler/response"
	"mygram-api/src/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service socialmedia.Service
}

func (h *Handler) CreateSocialMediaHandler(c *gin.Context) {
	request := request.Request{}
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
	result, err := h.service.CreateSocialMedia(request.MapToRecord(userData.ID))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, response.MapToPostResponse(*result))
}

func (h *Handler) GetAllSocialMediasHandler(c *gin.Context) {
	result, err := h.service.GetAllSocialMedias()
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.MapToBatchSocialMediaResponse(result))
}

func (h *Handler) UpdateSocialMediaHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	request := request.Request{}
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
	result, err := h.service.UpdateSocialMedia(id, request.MapToRecord(userData.ID))
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

func (h *Handler) DeleteSocialMediaHandler(c *gin.Context) {
	userData := helper.GetUserData(c)
	id, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteSocialMedia(id, userData.ID); err != nil {
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

	helper.CreateMessageResponse(c, http.StatusOK, "Your social media has been successfully deleted")
}

func NewHandler(service socialmedia.Service) *Handler {
	return &Handler{service}
}
