package response

import (
	"mygram-api/src/app/user/repository/record"
	"time"
)

type RegisterResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UpdateResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapToRegisterResponse(data record.User) RegisterResponse {
	return RegisterResponse{
		ID:       data.ID,
		Email:    data.Email,
		Username: data.Username,
		Age:      data.Age,
	}
}

func MapToUpdateResponse(data record.User) UpdateResponse {
	return UpdateResponse{
		ID:        data.ID,
		Email:     data.Email,
		Username:  data.Username,
		Age:       data.Age,
		UpdatedAt: data.UpdatedAt,
	}
}
