package request

type PostRequest struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
}

type UpdateRequest struct {
	Message string `json:"message"`
}
