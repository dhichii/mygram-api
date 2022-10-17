package socialmedia

import "mygram-api/src/app/socialmedia/repository/record"

type Service interface {
	CreateSocialMedia(socialMedia *record.SocialMedia) (*record.SocialMedia, error)
	GetAllSocialMedias() ([]record.SocialMedia, error)
	UpdateSocialMedia(id int, socialMedia *record.SocialMedia) (*record.SocialMedia, error)
	DeleteSocialMedia(id int) error
}

type Repository interface {
	CreateData(data *record.SocialMedia) (*record.SocialMedia, error)
	GetAllData() ([]record.SocialMedia, error)
	UpdateData(id int, data *record.SocialMedia) (*record.SocialMedia, error)
	DeleteData(id int) error
}
