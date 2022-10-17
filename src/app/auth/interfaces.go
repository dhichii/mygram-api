package auth

type Service interface {
	GetUserIdByPhotoId(photoId int) (int, error)
	GetUserIdByCommentId(commentId int) (int, error)
	GetUserIdBySocialMediaId(socialMediaId int) (int, error)
}

type Repository interface {
	GetUserIdByFeatureId(feature string, id int) (int, error)
}
