package request

import (
	"mygram-api/src/app/user/repository/record"
	"mygram-api/src/helper"
)

type RegisterRequest struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
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
