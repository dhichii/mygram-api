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

func (repo *repository) UpdateData(id int, data *record.Comment) (*record.Comment, error) {
	if err := repo.db.Where("id", id).Updates(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *repository) GetDataByID(id int) (*record.Comment, error) {
	record := &record.Comment{}
	if err := repo.db.First(record, "id", id).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (repo *repository) DeleteData(id int) error {
	return repo.db.Delete(new(record.Photo), "id", id).Error
}

func (repo *repository) GetUserIDByID(id int) (int, error) {
	userId := 0
	if err := repo.db.Model(new(record.Comment)).Select("user_id").First(&userId, "id", id).Error; err != nil {
		return 0, err
	}

	return userId, nil
}

func NewGORMRepository(db *gorm.DB) comment.Repository {
	return &repository{db}
}
