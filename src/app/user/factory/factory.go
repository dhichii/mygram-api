package factory

import (
	"mygram-api/src/app/user/handler"
	"mygram-api/src/app/user/repository"
	"mygram-api/src/app/user/service"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) handler.Handler {
	repo := repository.NewGORMRepository(conn)
	serv := service.NewService(repo)
	return *handler.NewHandler(serv)
}
