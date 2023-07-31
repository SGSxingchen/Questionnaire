package main

import (
	"github.com/gin-gonic/gin"
	"questionnaire/config"
	"questionnaire/db"
	"questionnaire/router"
)

func main() {
	config.InitConfig()
	db.InitDB()
	router.InitRouter()
	if !config.Config.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Router.Run(":" + config.Config.Server.Port)
}
