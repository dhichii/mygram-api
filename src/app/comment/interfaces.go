package comment

import "mygram-api/src/app/comment/repository/record"

type Service interface {
	CreateComment(comment *record.Comment) (*record.Comment, error)
	GetAllComments() ([]record.Comment, error)
	UpdateComment(id, userId int, message string) (*record.Comment, error)
	DeleteComment(id, userId int) error
}

type Repository interface {
	CreateData(data *record.Comment) (*record.Comment, error)
	GetAllData() ([]record.Comment, error)
	UpdateData(id int, data *record.Comment) (*record.Comment, error)
	GetDataByID(id int) (*record.Comment, error)
	DeleteData(id int) error
	GetUserIDByID(id int) (int, error)
}
