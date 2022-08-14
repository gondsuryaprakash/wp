package rfhandler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"waetherproject.com/core/domain"
	service "waetherproject.com/core/service/rfservice"
	"waetherproject.com/util"
	"waetherproject.com/util/logger"
	"waetherproject.com/util/notification"
	"waetherproject.com/utilities"
)

func GetRFById(svc service.RFService) gin.HandlerFunc {
	funcName := "rfhandler.GetRFById"
	logger.I(funcName)
	var exec context.Context
	return func(c *gin.Context) {
		exec = util.SetContext(c)
		id := c.Param("id")
		res := svc.GetRFById(&exec, cast.ToInt(id))
		c.JSON(cast.ToInt(utilities.RF_CODE_200), res)
	}

}

func AddRF(svc service.RFService) gin.HandlerFunc {
	funcName := "rfhandler.AddRF"
	logger.I(funcName)
	var exec context.Context
	return func(c *gin.Context) {
		exec = util.SetContext(c)
		var rfReqData domain.RfDataDomainReq
		err := c.Bind(&rfReqData)
		if err != nil {
			logger.E(funcName, err)
			response := notification.NewResponse(utilities.RF_CODE_500, utilities.GetMessageWithStatusCode(utilities.RF_CODE_500))
			c.JSON(200, response)
			return
		}

		response := svc.AddRF(&exec, &rfReqData)
		c.JSON(200, response)

	}

}

func UpdateRF(svc service.RFService) gin.HandlerFunc {
	funcName := "rfhandler.UpdateRF"
	logger.I(funcName)
	var exec context.Context
	return func(c *gin.Context) {
		exec = util.SetContext(c)
		var m domain.RfDataDomainReqForUpdate
		if err := c.Bind(&m); err != nil {
			logger.E(funcName, err)
			response := notification.NewResponse(utilities.RF_CODE_500, utilities.GetMessageWithStatusCode(utilities.RF_CODE_500))
			c.JSON(200, response)
			return
		}
		response := svc.UpdateRF(&exec, &m)
		c.JSON(200, response)
	}

}

func DeleteRF(svc service.RFService) gin.HandlerFunc {
	funcName := "rfhandler.DeleteRF"
	logger.I(funcName)
	var exec context.Context
	return func(c *gin.Context) {
		exec = util.SetContext(c)
		id := c.Param("id")
		response := svc.DeleteRF(&exec, cast.ToInt(id))
		c.JSON(200, response)
	}
}
