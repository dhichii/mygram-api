package user

import (
	"mygram-api/src/app/user/handler/request"
	"mygram-api/src/app/user/repository/record"
)

type Service interface {
	RegisterUser(user *record.User) (*record.User, error)
	LoginUser(login request.LoginRequest) (*string, error)
	UpdateUser(id string, user *record.User) (*record.User, error)
	DeleteUser(id string) error
}

type Repository interface {
	CreateData(data *record.User) (*record.User, error)
	FindDataByEmail(email string) (*record.User, error)
	UpdateData(id string, data *record.User) (*record.User, error)
	DeleteData(id string) error
}
