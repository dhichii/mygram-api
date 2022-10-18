package service

import (
	"mygram-api/src/app/comment"
	"mygram-api/src/app/comment/repository/record"
)

type service struct {
	repo comment.Repository
}

func (s *service) CreateComment(comment *record.Comment) (*record.Comment, error) {
	return s.repo.CreateData(comment)
}

func (s *service) GetAllComments() ([]record.Comment, error) {
	return s.repo.GetAllData()
}

func (s *service) UpdateComment(id int, message string) (*record.Comment, error) {
	return s.repo.UpdateData(id, message)
}

func (s *service) DeleteComment(id int) error {
	return s.repo.DeleteData(id)
}

func NewService(repo comment.Repository) comment.Service {
	return &service{repo}
}
