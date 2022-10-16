package factory

import (
	"mygram-api/src/app/comment/handler"
	"mygram-api/src/app/comment/repository"
	"mygram-api/src/app/comment/service"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) handler.Handler {
	repo := repository.NewGORMRepository(conn)
	serv := service.NewService(repo)
	return *handler.NewHandler(serv)
}
