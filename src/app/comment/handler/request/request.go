package request

import "mygram-api/src/app/comment/repository/record"

type PostRequest struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id" binding:"required"`
}

type UpdateRequest struct {
	Message string `json:"message" binding:"required"`
}

func (req *PostRequest) MapToRecord(userId int) *record.Comment {
	return &record.Comment{
		UserID:  userId,
		Message: req.Message,
		PhotoID: req.PhotoID,
	}
}
