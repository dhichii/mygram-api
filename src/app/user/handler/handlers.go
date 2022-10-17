package handler

import (
	"errors"
	"mygram-api/src/app/user"
	"mygram-api/src/app/user/handler/request"
	"mygram-api/src/app/user/handler/response"
	"mygram-api/src/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service user.Service
}

// @Tags User
// @Summary Register user account
// @ID register-user
// @Accept json
// @Produce json
// @Param RequestBody body request.RegisterRequest true "json request body"
// @Success 201 {object} response.RegisterResponse
// @Router /users/register [post]
func (h *Handler) RegisterUserHandler(c *gin.Context) {
	request := request.RegisterRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, helper.ValidateRequest(verr))
			return
		}

		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.RegisterUser(request.MapToRegisterRequest())
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			if strings.Contains(err.Error(), "email") {
				helper.CreateMessageResponse(c, http.StatusBadRequest,
					"email is already used")
				return
			}

			helper.CreateMessageResponse(c, http.StatusBadRequest,
				"username is already used")
			return
		}

		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, response.MapToRegisterResponse(*result))
}

// @Tags User
// @Summary Login user
// @ID login-user
// @Accept json
// @Produce json
// @Param RequestBody body request.LoginRequest true "json request body"
// @Success 200 {object} response.LoginResponse
// @Router /users/login [post]
func (h *Handler) LoginUserHandler(c *gin.Context) {
	request := request.LoginRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, helper.ValidateRequest(verr))
			return
		}

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

// @Tags User
// @Summary Update user
// @ID update-user
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body request.UpdateRequest true "json request body"
// @Success 200 {object} response.UpdateResponse
// @Router /users [put]
func (h *Handler) UpdateUserHandler(c *gin.Context) {
	request := request.UpdateRequest{}
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
	result, err := h.service.UpdateUser(userData.ID, request.MapToUpdateRequest())
	if err != nil {
		if err.Error() == helper.NOTFOUND {
			helper.CreateMessageResponse(c, http.StatusNotFound,
				http.StatusText(http.StatusNotFound))
			return
		}

		if strings.Contains(err.Error(), "unique") {
			if strings.Contains(err.Error(), "email") {
				helper.CreateMessageResponse(c, http.StatusBadRequest,
					"email is already used")
				return
			}

			helper.CreateMessageResponse(c, http.StatusBadRequest,
				"username is already used")
			return
		}

		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.MapToUpdateResponse(*result))
}

// @Tags User
// @Summary Delete user account
// @ID delete-user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} structs.Message
// @Router /users [delete]
func (h *Handler) DeleteUserHandler(c *gin.Context) {
	userData := helper.GetUserData(c)
	if err := h.service.DeleteUser(userData.ID); err != nil {
		if err.Error() == helper.NOTFOUND {
			helper.CreateMessageResponse(c, http.StatusNotFound,
				http.StatusText(http.StatusNotFound))
			return
		}

		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	helper.CreateMessageResponse(c, http.StatusOK, "Your account has been successfully deleted")
}

func NewHandler(service user.Service) *Handler {
	return &Handler{service}
}
