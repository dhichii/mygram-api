package repository

import (
	"fmt"
	"mygram-api/src/app/auth"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// GetUserIDByID return the UserID by the given Photo ID
func (repo *repository) GetUserIdByFeatureId(feature string, id int) (int, error) {
	userId := 0
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE id=%d", feature, id)
	if err := repo.db.Raw(query).Scan(&userId).Error; err != nil {
		return 0, err
	}

	return userId, nil
}

func NewGORMRepository(db *gorm.DB) auth.Repository {
	return &repository{db}
}
