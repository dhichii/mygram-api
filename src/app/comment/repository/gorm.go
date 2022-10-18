package repository

import (
	"mygram-api/src/app/comment"
	"mygram-api/src/app/comment/repository/record"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (repo *repository) CreateData(data *record.Comment) (*record.Comment, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *repository) GetAllData() ([]record.Comment, error) {
	records := []record.Comment{}
	if err := repo.db.Preload("User").Preload("Photo").Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (repo *repository) UpdateData(id int, message string) (*record.Comment, error) {
	data := &record.Comment{}
	if err := repo.db.Model(&data).Where("id", id).Update("message", message).
		Scan(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *repository) DeleteData(id int) error {
	return repo.db.Delete(new(record.Photo), "id", id).Error
}

func NewGORMRepository(db *gorm.DB) comment.Repository {
	return &repository{db}
}
