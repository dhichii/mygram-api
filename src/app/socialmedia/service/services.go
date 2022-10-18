package service

import (
	"mygram-api/src/app/socialmedia"
	"mygram-api/src/app/socialmedia/repository/record"
	"mygram-api/src/helper/errs"
)

type service struct {
	repository socialmedia.Repository
}

func (s *service) CreateSocialMedia(socialMedia *record.SocialMedia) (*record.SocialMedia, errs.MessageErr) {
	return s.repository.CreateData(socialMedia)
}

func (s *service) GetAllSocialMedias() ([]record.SocialMedia, errs.MessageErr) {
	return s.repository.GetAllData()
}

func (s *service) UpdateSocialMedia(id int, socialMedia *record.SocialMedia) (*record.SocialMedia, errs.MessageErr) {
	return s.repository.UpdateData(id, socialMedia)
}

func (s *service) DeleteSocialMedia(id int) errs.MessageErr {
	return s.repository.DeleteData(id)
}

func NewService(repo socialmedia.Repository) socialmedia.Service {
	return &service{repo}
}
