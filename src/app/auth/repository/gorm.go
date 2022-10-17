package repository

import (
	"errors"
	"fmt"
	"mygram-api/src/app/auth"
	"mygram-api/src/helper"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// GetUserIDByFeatureID return the UserID by the given feature ID *example: photos, comments, etc
func (repo *repository) GetUserIDByFeatureID(feature string, id int) (int, error) {
	userId := 0
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE id=%d", feature, id)
	err := repo.db.Raw(query).Scan(&userId).Error
	if err != nil {
		return 0, err
	}

	if err == nil && userId == 0 {
		return 0, errors.New(helper.NOTFOUND)
	}

	return userId, nil
}

func NewGORMRepository(db *gorm.DB) auth.Repository {
	return &repository{db}
}
