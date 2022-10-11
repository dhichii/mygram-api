package service

import (
	"errors"
	"mygram-api/src/app/photo"
	"mygram-api/src/app/photo/repository/record"
	"mygram-api/src/helper"
)

type service struct {
	repo photo.Repository
}

// PostPhoto post new photo
func (s *service) PostPhoto(photo *record.Photo) (*record.Photo, error) {
	return s.repo.CreateData(photo)
}

// GetAllPhotos get all photos
func (s *service) GetAllPhotos() ([]record.Photo, error) {
	return s.repo.GetAllData()
}

// UpdatePhoto update photo data by the given id
// it will be return forbidden error if the UserID is different
func (s *service) UpdatePhoto(id int, photo *record.Photo) (*record.Photo, error) {
	userId, err := s.repo.GetUserIDByID(id)
	if err != nil {
		return nil, err
	}

	if userId != photo.UserID {
		return nil, errors.New(helper.FORBIDDEN)
	}

	return s.repo.UpdateData(id, photo)
}

// DeletePhoto delete the photo by the given id
// it will be return forbidden error if the UserID is different
func (s *service) DeletePhoto(id, userId int) error {
	userIdResult, err := s.repo.GetUserIDByID(id)
	if err != nil {
		return err
	}

	if userIdResult != userId {
		return errors.New(helper.FORBIDDEN)
	}

	return s.repo.DeleteData(id)
}

func NewService(repo photo.Repository) photo.Service {
	return &service{repo}
}
