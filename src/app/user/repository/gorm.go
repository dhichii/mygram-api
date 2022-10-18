package repository

import (
	"mygram-api/src/app/user"
	"mygram-api/src/app/user/repository/record"
	"mygram-api/src/helper"
	"mygram-api/src/helper/errs"
	"net/http"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// CreateData create new data from the given input
func (repo *repository) CreateData(data *record.User) (*record.User, errs.MessageErr) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, errs.NewCustomError(http.StatusInternalServerError, err.Error())
	}

	return data, nil
}

// FindDataByEmail find the id, email, and password by the given email
func (repo *repository) FindDataByEmail(email string) (*record.User, errs.MessageErr) {
	record := &record.User{}
	if err := repo.db.Select("id", "email", "password").First(record, "email", email).
		Error; err != nil {
		if err.Error() == helper.NOTFOUND {
			return nil, errs.NewCustomError(http.StatusUnauthorized, "invalid email or password")
		}

		return nil, errs.NewError(http.StatusInternalServerError)
	}

	return record, nil
}

// UpdateData update the data by the given id
func (repo *repository) UpdateData(id int, data *record.User) (*record.User, errs.MessageErr) {
	if err := repo.db.Where("id", id).Updates(data).Scan(data).Error; err != nil {
		return nil, errs.NewCustomError(http.StatusInternalServerError, err.Error())
	}

	return data, nil
}

// DeleteData delete the data by the given id
func (repo *repository) DeleteData(id int) errs.MessageErr {
	if err := repo.db.Delete(new(record.User), "id", id).Error; err != nil {
		return errs.NewError(http.StatusInternalServerError)
	}

	return nil
}

func NewGORMRepository(db *gorm.DB) user.Repository {
	return &repository{db}
}
