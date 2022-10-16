package service

import (
	"errors"
	"mygram-api/src/app/comment"
	"mygram-api/src/app/comment/repository/record"
	"mygram-api/src/helper"
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

func (s *service) UpdateComment(id, userId int, message string) (*record.Comment, error) {
	record, err := s.repo.GetDataByID(id)
	if err != nil {
		return nil, err
	}

	if record.UserID != userId {
		return nil, errors.New(helper.FORBIDDEN)
	}

	record.Message = message
	return s.repo.UpdateData(id, record)
}

func (s *service) DeleteComment(id, userId int) error {
	userIdResult, err := s.repo.GetUserIDByID(id)
	if err != nil {
		return err
	}

	if userIdResult != userId {
		return errors.New(helper.FORBIDDEN)
	}

	return s.repo.DeleteData(id)
}

func NewService(repo comment.Repository) comment.Service {
	return &service{repo}
}
