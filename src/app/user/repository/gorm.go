package repository

import (
	"errors"
	"mygram-api/src/app/user"
	"mygram-api/src/app/user/repository/record"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// CreateData create new data from the given input
func (repo *repository) CreateData(data *record.User) (*record.User, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// FindDataByEmail find the id, email, and password by the given email
func (repo *repository) FindDataByEmail(email string) (*record.User, error) {
	record := &record.User{}
	if err := repo.db.Select("id", "email", "password").
		First(record, "email", email).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// UpdateData update the data by the given id
func (repo *repository) UpdateData(id int, data *record.User) (*record.User, error) {
	query := repo.db.Where("id", id).Updates(data)
	err := query.Error
	if err == nil && query.RowsAffected < 1 {
		return nil, errors.New("record not found")
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

// DeleteData delete the data by the given id
func (repo *repository) DeleteData(id int) error {
	query := repo.db.Delete(new(record.User), "id", id)
	if query.Error == nil && query.RowsAffected < 1 {
		return errors.New("record not found")
	}

	return query.Error
}

func NewGORMRepository(db *gorm.DB) user.Repository {
	return &repository{db}
}