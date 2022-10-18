package comment

import "mygram-api/src/app/comment/repository/record"

type Service interface {
	CreateComment(comment *record.Comment) (*record.Comment, error)
	GetAllComments() ([]record.Comment, error)
	UpdateComment(id int, message string) (*record.Comment, error)
	DeleteComment(id int) error
}

type Repository interface {
	CreateData(data *record.Comment) (*record.Comment, error)
	GetAllData() ([]record.Comment, error)
	UpdateData(id int, data string) (*record.Comment, error)
	DeleteData(id int) error
}
