package service

import (
	"errors"
	"mygram-api/src/app/socialmedia"
	"mygram-api/src/app/socialmedia/repository/record"
	"mygram-api/src/helper"
)

type service struct {
	repository socialmedia.Repository
}

func (s *service) CreateSocialMedia(socialMedia *record.SocialMedia) (*record.SocialMedia, error) {
	return s.repository.CreateData(socialMedia)
}

func (s *service) GetAllSocialMedias() ([]record.SocialMedia, error) {
	return s.repository.GetAllData()
}

func (s *service) UpdateSocialMedia(id int, socialMedia *record.SocialMedia) (*record.SocialMedia, error) {
	userId, err := s.repository.GetUserIDByID(id)
	if err != nil {
		return nil, err
	}

	if userId != socialMedia.UserID {
		return nil, errors.New(helper.FORBIDDEN)
	}

	return s.repository.UpdateData(id, socialMedia)
}

func (s *service) DeleteSocialMedia(id int, userId int) error {
	userIdResult, err := s.repository.GetUserIDByID(id)
	if err != nil {
		return err
	}

	if userIdResult != userId {
		return errors.New(helper.FORBIDDEN)
	}

	return s.repository.DeleteData(id)
}

func NewService(repo socialmedia.Repository) socialmedia.Service {
	return &service{repo}
}
