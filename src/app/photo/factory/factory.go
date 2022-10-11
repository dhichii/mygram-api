package factory

import (
	"mygram-api/src/app/photo/handler"
	"mygram-api/src/app/photo/repository"
	"mygram-api/src/app/photo/service"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) handler.Handler {
	repo := repository.NewGORMRepository(conn)
	serv := service.NewService(repo)
	return *handler.NewHandler(serv)
}
