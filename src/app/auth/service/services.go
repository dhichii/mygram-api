package service

import "mygram-api/src/app/auth"

type service struct {
	repo auth.Repository
}

func (s *service) GetUserIdByPhotoId(photoId int) (int, error) {
	return s.repo.GetUserIdByFeatureId("photos", photoId)
}

func (s *service) GetUserIdByCommentId(commentId int) (int, error) {
	return s.repo.GetUserIdByFeatureId("comments", commentId)
}

func (s *service) GetUserIdBySocialMediaId(socialMediaId int) (int, error) {
	return s.repo.GetUserIdByFeatureId("social_medias", socialMediaId)
}

func NewService(repo auth.Repository) auth.Service {
	return &service{repo}
}
