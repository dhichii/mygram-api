package service

import (
	"mygram-api/src/app/photo"
	"mygram-api/src/app/photo/repository/record"
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
	return s.repo.UpdateData(id, photo)
}

// DeletePhoto delete the photo by the given id
// it will be return forbidden error if the UserID is different
func (s *service) DeletePhoto(id int) error {
	return s.repo.DeleteData(id)
}

func NewService(repo photo.Repository) photo.Service {
	return &service{repo}
}
