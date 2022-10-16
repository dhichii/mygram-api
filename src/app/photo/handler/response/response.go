package response

import (
	"mygram-api/src/app/photo/repository/record"
	"time"
)

type PostResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
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
	User      UserResponse `json:"User"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapToPostResponse(res record.Photo) PostResponse {
	return PostResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		CreatedAt: res.CreatedAt,
	}
}

func MapToBatchPhotoResponse(res []record.Photo) (responses []PhotoResponse) {
	for _, photo := range res {
		responses = append(responses, MapToPhotoResponse(photo))
	}
	return
}

func MapToPhotoResponse(res record.Photo) PhotoResponse {
	return PhotoResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		User: UserResponse{
			Email:    res.User.Email,
			Username: res.User.Username,
		},
	}
}

func MapToUpdateResponse(res record.Photo) UpdateResponse {
	return UpdateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		UpdatedAt: res.UpdatedAt,
	}
}
