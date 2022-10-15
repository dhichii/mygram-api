package factory

import (
	"mygram-api/src/app/socialmedia/handler"
	"mygram-api/src/app/socialmedia/repository"
	"mygram-api/src/app/socialmedia/service"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) handler.Handler {
	repo := repository.NewGORMRepository(conn)
	serv := service.NewService(repo)
	return *handler.NewHandler(serv)
}
