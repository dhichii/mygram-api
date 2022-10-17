package service

import "mygram-api/src/app/auth"

type service struct {
	repo auth.Repository
}

func (s *service) GetUserIDByPhotoID(photoId int) (int, error) {
	return s.repo.GetUserIDByFeatureID("photos", photoId)
}

func (s *service) GetUserIDByCommentID(commentId int) (int, error) {
	return s.repo.GetUserIDByFeatureID("comments", commentId)
}

func (s *service) GetUserIDBySocialMediaID(socialMediaId int) (int, error) {
	return s.repo.GetUserIDByFeatureID("social_medias", socialMediaId)
}

func NewService(repo auth.Repository) auth.Service {
	return &service{repo}
}
