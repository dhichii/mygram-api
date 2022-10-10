package adapter

import (
	userFactory "mygram-api/src/app/user/factory"
	userHandler "mygram-api/src/app/user/handler"
	"mygram-api/src/database"
)

type handlers struct {
	User userHandler.Handler
}

func Init() handlers {
	conn := database.GetPostgresConnection()
	return handlers{
		User: userFactory.Factory(conn),
	}
}
