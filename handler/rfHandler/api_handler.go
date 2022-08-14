package rfhandler

import (
	"github.com/gin-gonic/gin"

	"waetherproject.com/core/persistence"
	rfservice "waetherproject.com/core/service/rfservice"
	"waetherproject.com/handler/rfHandler/rfhandler"
	"waetherproject.com/util/logger"
)

func SetRoutes(router *gin.Engine) {

	funcName := "handler.RfHandler"
	logger.I(funcName)

	rfPersistence := persistence.NewRfPersistence()

	rfSvc := rfservice.NewRFService(rfPersistence)

	rfGroup := router.Group("/rf")
	{
		rfGroup.GET(":id", rfhandler.GetRFById(rfSvc))
		rfGroup.POST("add", rfhandler.AddRF(rfSvc))
		rfGroup.PUT("update", rfhandler.UpdateRF(rfSvc))
		rfGroup.DELETE("delete/:id", rfhandler.DeleteRF(rfSvc))

	}

}
