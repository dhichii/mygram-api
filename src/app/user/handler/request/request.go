package request

import (
	"mygram-api/src/app/user/repository/record"
	"mygram-api/src/helper"
)

type RegisterRequest struct {
	Age      int    `json:"age" binding:"required,min=9"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
}

func (req *RegisterRequest) MapToRegisterRequest() *record.User {
	return &record.User{
		Age:      req.Age,
		Email:    req.Email,
		Password: helper.CreateHash(req.Password),
		Username: req.Username,
	}
}

func (req *UpdateRequest) MapToUpdateRequest() *record.User {
	return &record.User{
		Email:    req.Email,
		Username: req.Username,
	}
}
