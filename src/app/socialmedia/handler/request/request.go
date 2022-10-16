package request

import "mygram-api/src/app/socialmedia/repository/record"

type Request struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}

func (r *Request) MapToRecord(userID int) *record.SocialMedia {
	return &record.SocialMedia{
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserID:         userID,
	}
}
