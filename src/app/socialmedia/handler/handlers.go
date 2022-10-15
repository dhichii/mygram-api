package handler

import (
	"mygram-api/src/app/socialmedia"
	"mygram-api/src/app/socialmedia/handler/request"
	"mygram-api/src/app/socialmedia/handler/response"
	"mygram-api/src/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service socialmedia.Service
}

func (h *Handler) CreateSocialMediaHandler(c *gin.Context) {
	request := request.Request{}
	userData := helper.GetUserData(c)

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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
	request := request.Request{}
	userData := helper.GetUserData(c)
	id, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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
