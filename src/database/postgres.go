package database

import (
	"fmt"
	"log"
	user "mygram-api/src/app/user/repository/record"
	"mygram-api/src/config/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresConn *gorm.DB

func InitPostgres() {
	var err error

	config := env.GetSQLEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port)
	postgresConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	postgresConn.AutoMigrate(&user.User{})
}

func GetPostgresConnection() *gorm.DB {
	return postgresConn
}
