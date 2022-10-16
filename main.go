package main

import (
	"mygram-api/src/config/env"
	"mygram-api/src/database"
	"mygram-api/src/router"
)

// @contact.name   Adhicitta Masran
// @contact.email  adhicittamasran@gmail.com
// @license.name  MIT
func main() {
	database.InitPostgres()
	router.StartServer().Run(":" + env.GetServerEnv())
}
