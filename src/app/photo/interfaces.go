package photo

import (
	"mygram-api/src/app/photo/repository/record"
	"mygram-api/src/helper/errs"
)

type Service interface {
	PostPhoto(photo *record.Photo) (*record.Photo, errs.MessageErr)
	GetAllPhotos() ([]record.Photo, errs.MessageErr)
	UpdatePhoto(id int, photo *record.Photo) (*record.Photo, errs.MessageErr)
	DeletePhoto(id int) errs.MessageErr
}

type Repository interface {
	CreateData(data *record.Photo) (*record.Photo, errs.MessageErr)
	GetAllData() ([]record.Photo, errs.MessageErr)
	UpdateData(id int, data *record.Photo) (*record.Photo, errs.MessageErr)
	DeleteData(id int) errs.MessageErr
}
