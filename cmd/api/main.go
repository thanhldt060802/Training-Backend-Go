package main

import (
	"training-project/internal/config"
	"training-project/route"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitConfig()
	db := config.ConnectDB()
	defer db.Close()

	r := gin.Default()
	route.RegisterRoutes(r, db)
	r.Run(":" + config.AppConfig.AppPort)

}
