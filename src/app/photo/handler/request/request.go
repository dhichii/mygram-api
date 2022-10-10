package request

import "mygram-api/src/app/photo/repository/record"

type Request struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

func (req *Request) MapToRecord(userId int) *record.Photo {
	return &record.Photo{
		UserID:   userId,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}
}
