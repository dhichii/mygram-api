package response

import "time"

type PostResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type Response struct {
	SocialMedias []SocialMediaResponse `json:"social_medias"`
}

type SocialMediaResponse struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	SocialMediaUrl string       `json:"social_media_url"`
	UserID         int          `json:"user_id"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	User           UserResponse `json:"User"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UpdateResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
