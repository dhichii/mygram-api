package adapter

import (
	photo_factory "mygram-api/src/app/photo/factory"
	photo_handler "mygram-api/src/app/photo/handler"
	user_factory "mygram-api/src/app/user/factory"
	user_handler "mygram-api/src/app/user/handler"
	"mygram-api/src/database"
)

type handlers struct {
	User  user_handler.Handler
	Photo photo_handler.Handler
}

func Init() handlers {
	conn := database.GetPostgresConnection()
	return handlers{
		User:  user_factory.Factory(conn),
		Photo: photo_factory.Factory(conn),
	}
}
