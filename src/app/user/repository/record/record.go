package record

import "time"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar"`
	Email     string `gorm:"type:varchar"`
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
