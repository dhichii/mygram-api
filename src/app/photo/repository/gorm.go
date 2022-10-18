package repository

import (
	"errors"
	"mygram-api/src/app/photo"
	"mygram-api/src/app/photo/repository/record"
	"mygram-api/src/helper"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// CreateData create new data from the given input
func (repo *repository) CreateData(data *record.Photo) (*record.Photo, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// GetAllData find all data
func (repo *repository) GetAllData() ([]record.Photo, error) {
	records := []record.Photo{}
	if err := repo.db.Preload("User").Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

// UpdateData Update the data by the given id
func (repo *repository) UpdateData(id int, data *record.Photo) (*record.Photo, error) {
	query := repo.db.Where("id", id).Updates(data).Scan(data)
	err := query.Error
	if err != nil {
		return nil, err
	}

	if err == nil && query.RowsAffected < 1 {
		return nil, errors.New(helper.NOTFOUND)
	}

	data.ID = id
	return data, nil
}

// DeleteData delete the data by the given id
func (repo *repository) DeleteData(id int) error {
	query := repo.db.Delete(new(record.Photo), "id", id)
	if query.Error == nil && query.RowsAffected < 1 {
		return errors.New(helper.NOTFOUND)
	}

	return query.Error
}

func NewGORMRepository(db *gorm.DB) photo.Repository {
	return &repository{db}
}
