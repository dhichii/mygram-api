package request

import "mygram-api/src/app/socialmedia/repository/record"

type Request struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

func (r *Request) MapToRecord(userID int) *record.SocialMedia {
	return &record.SocialMedia{
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserID:         userID,
	}
}
