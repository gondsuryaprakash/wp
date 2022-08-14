package main

import (
	"github.com/gin-gonic/gin"
	rfhandler "waetherproject.com/handler/rfHandler"
	"waetherproject.com/util"
	"waetherproject.com/util/logger"
)

func init() {
	logger.I("main.init")
	util.Connection()
}

func main() {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(gin.Logger())
	rfhandler.SetRoutes(app)

	app.Run(":" + util.GetConfigValue("PORT"))

}
