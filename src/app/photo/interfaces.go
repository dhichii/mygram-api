package photo

import "mygram-api/src/app/photo/repository/record"

type Service interface {
	PostPhoto(photo *record.Photo) (*record.Photo, error)
	GetAllPhotos() ([]record.Photo, error)
	UpdatePhoto(id int, photo *record.Photo) (*record.Photo, error)
	DeletePhoto(id int) error
}

type Repository interface {
	CreateData(data *record.Photo) (*record.Photo, error)
	GetAllData() ([]record.Photo, error)
	UpdateData(id int, data *record.Photo) (*record.Photo, error)
	DeleteData(id int) error
}
