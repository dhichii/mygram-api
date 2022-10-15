package adapter

import (
	photo_factory "mygram-api/src/app/photo/factory"
	photo_handler "mygram-api/src/app/photo/handler"
	social_media_factory "mygram-api/src/app/socialmedia/factory"
	social_media_handler "mygram-api/src/app/socialmedia/handler"
	user_factory "mygram-api/src/app/user/factory"
	user_handler "mygram-api/src/app/user/handler"
	"mygram-api/src/database"
)

type handlers struct {
	User        user_handler.Handler
	Photo       photo_handler.Handler
	SocialMedia social_media_handler.Handler
}

func Init() handlers {
	conn := database.GetPostgresConnection()
	return handlers{
		User:        user_factory.Factory(conn),
		Photo:       photo_factory.Factory(conn),
		SocialMedia: social_media_factory.Factory(conn),
	}
}
