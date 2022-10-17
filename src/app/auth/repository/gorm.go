package repository

import (
	"mygram-api/src/app/auth"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// GetUserIDByID return the UserID by the given Photo ID
func (repo *repository) GetUserIdByFeatureId(feature string, id int) (int, error) {
	userId := 0
	if err := repo.db.Raw("SELECT user_id FROM ? WHERE id=?", feature, id).Scan(&userId).Error; err != nil {
		return 0, err
	}

	return userId, nil
}

func NewGORMRepository(db *gorm.DB) auth.Repository {
	return &repository{db}
}
