package response

import "time"

type PostResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"string"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoResponse struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserID    int          `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"string"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
