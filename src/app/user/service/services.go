package service

import (
	"errors"
	"mygram-api/src/app/user"
	"mygram-api/src/app/user/handler/request"
	"mygram-api/src/app/user/repository/record"
	"mygram-api/src/helper"
)

type service struct {
	repo user.Repository
}

// RegisterUser register the new user by create new data
func (s *service) RegisterUser(user *record.User) (*record.User, error) {
	return s.repo.CreateData(user)
}

// LoginUser login the user and return the jwt if request valid
func (s *service) LoginUser(login request.LoginRequest) (*string, error) {
	record, err := s.repo.FindDataByEmail(login.Email)
	if err != nil {
		if err.Error() == helper.NOTFOUND {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !helper.ValidateHash(login.Password, record.Password) {
		return nil, errors.New("invalid email or password")
	}

	token := helper.GenerateJWT(record.ID)
	return &token, nil
}

// UpdateUser update user data by the given id
func (s *service) UpdateUser(id int, user *record.User) (*record.User, error) {
	return s.repo.UpdateData(id, user)
}

// DeleteUser delete the user data by the given id
func (s *service) DeleteUser(id int) error {
	return s.repo.DeleteData(id)
}

func NewService(repo user.Repository) user.Service {
	return &service{repo}
}
