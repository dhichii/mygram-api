package repository

import (
	"errors"
	"mygram-api/src/app/socialmedia"
	"mygram-api/src/app/socialmedia/repository/record"
	"mygram-api/src/helper"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (repo *repository) CreateData(data *record.SocialMedia) (*record.SocialMedia, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *repository) GetAllData() ([]record.SocialMedia, error) {
	records := []record.SocialMedia{}
	if err := repo.db.Preload("User").Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (repo *repository) UpdateData(id int, data *record.SocialMedia) (*record.SocialMedia, error) {
	query := repo.db.Where("id", id).Updates(data)
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

func (repo *repository) DeleteData(id int) error {
	query := repo.db.Delete(new(record.SocialMedia), "id", id)
	if query.Error == nil && query.RowsAffected < 1 {
		return errors.New(helper.NOTFOUND)
	}

	return query.Error
}

func NewGORMRepository(db *gorm.DB) socialmedia.Repository {
	return &repository{db}
}
