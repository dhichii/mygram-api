package response

import (
	"mygram-api/src/app/socialmedia/repository/record"
	"time"
)

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

func MapToPostResponse(res record.SocialMedia) PostResponse {
	return PostResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		CreatedAt:      res.CreatedAt,
	}
}

func MapToBatchSocialMediaResponse(res []record.SocialMedia) Response {
	socialMedias := []SocialMediaResponse{}
	for _, socialMedia := range res {
		socialMedias = append(socialMedias, MapToSocialMediaResponse(socialMedia))
	}

	return Response{SocialMedias: socialMedias}
}

func MapToSocialMediaResponse(res record.SocialMedia) SocialMediaResponse {
	return SocialMediaResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
		User: UserResponse{
			ID:       res.UserID,
			Username: res.User.Username,
		},
	}
}

func MapToUpdateResponse(res record.SocialMedia) UpdateResponse {
	return UpdateResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		UpdatedAt:      res.UpdatedAt,
	}
}
