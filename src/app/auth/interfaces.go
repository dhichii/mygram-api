package auth

type Service interface {
	GetUserIDByPhotoID(photoId int) (int, error)
	GetUserIDByCommentID(commentId int) (int, error)
	GetUserIDBySocialMediaID(socialMediaId int) (int, error)
}

type Repository interface {
	GetUserIDByFeatureID(feature string, id int) (int, error)
}
