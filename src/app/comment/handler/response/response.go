package response

import (
	"mygram-api/src/app/comment/repository/record"
	"time"
)

type PostResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponse struct {
	ID        int           `json:"id"`
	Message   string        `json:"message"`
	PhotoID   int           `json:"photo_id"`
	UserID    int           `json:"user_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      UserResponse  `json:"User"`
	Photo     PhotoResponse `json:"Photo"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type UpdateResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapToPostResponse(res record.Comment) PostResponse {
	return PostResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoID:   res.PhotoID,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
	}
}

func MapToBatchCommentResponse(res []record.Comment) (responses []CommentResponse) {
	for _, comment := range res {
		responses = append(responses, mapToCommentResponse(comment))
	}
	return
}

func mapToCommentResponse(res record.Comment) CommentResponse {
	return CommentResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoID:   res.PhotoID,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		User: UserResponse{
			ID:       res.User.ID,
			Email:    res.User.Email,
			Username: res.User.Username,
		},
		Photo: PhotoResponse{
			ID:       res.Photo.ID,
			Title:    res.Photo.Title,
			Caption:  res.Photo.Caption,
			PhotoUrl: res.Photo.PhotoUrl,
			UserID:   res.Photo.UserID,
		},
	}
}

func MapToUpdateResponse(res record.Comment) UpdateResponse {
	return UpdateResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoID:   res.PhotoID,
		UserID:    res.UserID,
		UpdatedAt: res.UpdatedAt,
	}
}
