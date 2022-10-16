package handler

import (
	"errors"
	"mygram-api/src/app/comment"
	"mygram-api/src/app/comment/handler/request"
	"mygram-api/src/app/comment/handler/response"
	"mygram-api/src/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service comment.Service
}

func (h *Handler) CreateCommentHandler(c *gin.Context) {
	request := request.PostRequest{}
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
	result, err := h.service.CreateComment(request.MapToRecord(userData.ID))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, response.MapToPostResponse(*result))
}

func (h *Handler) GetAllCommentsHandler(c *gin.Context) {
	result, err := h.service.GetAllComments()
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.MapToBatchCommentResponse(result))
}

// func (h *Handler) GetAllCommentsHandler(c *gin.Context)
func (h *Handler) UpdateCommentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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
	result, err := h.service.UpdateComment(id, userData.ID, request.Message)
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

func (h *Handler) DeleteCommentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userData := helper.GetUserData(c)
	if err := h.service.DeleteComment(id, userData.ID); err != nil {
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

	helper.CreateMessageResponse(c, http.StatusOK, "Your comment has been successfully deleted")
}

func NewHandler(service comment.Service) *Handler {
	return &Handler{service}
}
