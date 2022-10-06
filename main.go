package main

import (
	"mygram-api/src/config/env"
	"mygram-api/src/database"
	"mygram-api/src/router"
)

func main() {
	database.InitPostgres()
	router.StartServer().Run(":" + env.GetServerEnv())
}
