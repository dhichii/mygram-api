package record

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;size:191"`
	Username  string    `gorm:"type:varchar"`
	Email     string    `gorm:"type:varchar"`
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
