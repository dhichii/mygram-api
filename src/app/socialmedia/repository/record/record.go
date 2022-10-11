package record

type SocialMedia struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar;not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserID         int
}

type User struct {
	ID       int `gorm:"primaryKey"`
	Username int
}
