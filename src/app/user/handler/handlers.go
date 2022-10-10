package handler

import (
	"mygram-api/src/app/user"
	"mygram-api/src/app/user/handler/request"
	"mygram-api/src/app/user/handler/response"
	"mygram-api/src/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service user.Service
}

func (h *Handler) RegisterUserHandler(c *gin.Context) {
	request := request.RegisterRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.RegisterUser(request.MapToRegisterRequest())
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, response.MapToRegisterResponse(*result))
}

func (h *Handler) LoginUserHandler(c *gin.Context) {
	request := request.LoginRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.LoginUser(request)
	if err != nil {
		if err.Error() == "invalid email or password" {
			helper.CreateMessageResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.LoginResponse{Token: *token})
}

func (h *Handler) UpdateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	request := request.UpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.UpdateUser(id, request.MapToUpdateRequest())
	if err != nil {
		if err.Error() == "record not found" {
			helper.CreateMessageResponse(c, http.StatusNotFound, err.Error())
			return
		}

		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.MapToUpdateResponse(*result))
}

func (h *Handler) DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		if err.Error() == "record not found" {
			helper.CreateMessageResponse(c, http.StatusNotFound, err.Error())
			return
		}

		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	helper.CreateMessageResponse(c, http.StatusOK, "your account has been successfully deleted")
}

func NewHandler(service user.Service) *Handler {
	return &Handler{service}
}
