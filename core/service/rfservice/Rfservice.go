package service

import (
	"context"

	"waetherproject.com/core/domain"
	"waetherproject.com/core/persistence"
	"waetherproject.com/util/logger"
	"waetherproject.com/util/notification"
	"waetherproject.com/utilities"
)

type RfSvc struct {
	RfPersistence persistence.RfPersistence
}

type RFService interface {
	GetRFById(ctx *context.Context, id int) notification.Response
	AddRF(ctx *context.Context, v *domain.RfDataDomainReq) notification.Response
	UpdateRF(ctx *context.Context, v *domain.RfDataDomainReqForUpdate) notification.Response
	DeleteRF(ctx *context.Context, id int) notification.Response
}

func NewRFService(rfPersistence persistence.RfPersistence) *RfSvc {
	return &RfSvc{
		RfPersistence: rfPersistence,
	}
}

func (c *RfSvc) GetRFById(ctx *context.Context, id int) notification.Response {
	funcName := "service.GetRFById"
	logger.I(funcName)
	v, err := c.RfPersistence.GetRFById(id)
	logger.I(funcName, 32)
	if err != nil {
		errCode := utilities.RF_CODE_500
		errMsg := utilities.GetMessageWithStatusCode(errCode)
		return notification.NewResponse(errCode, errMsg)
	}
	succCode := utilities.RF_CODE_200
	msg := utilities.GetMessageWithStatusCode(succCode)
	return notification.NewResponseWithModel(succCode, msg, v)
}

func (c *RfSvc) AddRF(ctx *context.Context, v *domain.RfDataDomainReq) notification.Response {
	m := &domain.RfDataDomain{
		Latitude:  v.Latitude,
		Longitude: v.Longitude,
		RF:        v.RF,
	}
	err := c.RfPersistence.AddRF(m)
	if err != nil {
		errCode := utilities.RF_CODE_500
		errMsg := utilities.GetMessageWithStatusCode(utilities.RF_CODE_500)
		return notification.NewResponse(errCode, errMsg)
	}
	succCode := utilities.RF_CODE_200
	msg := utilities.GetMessageWithStatusCode(succCode)
	return notification.NewResponseWithModel(succCode, msg, nil)

}

func (c *RfSvc) UpdateRF(ctx *context.Context, v *domain.RfDataDomainReqForUpdate) notification.Response {
	funcName := "rfservice.UpdateRF"
	logger.I(funcName)
	logger.D(funcName, v.Id)
	logger.D(funcName, v.Latitude)
	logger.D(funcName, v.Longitude)
	logger.D(funcName, v.RF)
	var colums []string

	if v.Latitude != nil {
		colums = append(colums, "Latitude")
	}
	if v.Longitude != nil {
		colums = append(colums, "Longitude")
	}
	if v.RF != nil {
		colums = append(colums, "RF")
	}
	m := &domain.RfDataDomain{
		Latitude:  v.Latitude,
		Longitude: v.Longitude,
		RF:        v.RF,
		Id:        v.Id,
	}

	if err := c.RfPersistence.UpdateRFById(m, colums...); err != nil {
		logger.E(funcName, err)
		errCode := utilities.RF_CODE_500
		errMsg := utilities.GetMessageWithStatusCode(utilities.RF_CODE_500)
		return notification.NewResponse(errCode, errMsg)
	}
	succCode := utilities.RF_CODE_200
	msg := utilities.GetMessageWithStatusCode(succCode)
	return notification.NewResponseWithModel(succCode, msg, nil)
}

func (c *RfSvc) DeleteRF(ctx *context.Context, id int) notification.Response {

	if err := c.RfPersistence.DeleteRF(id); err != nil {
		logger.E(err)
		errCode := utilities.RF_CODE_500
		errMsg := utilities.GetMessageWithStatusCode(utilities.RF_CODE_500)
		return notification.NewResponse(errCode, errMsg)
	}
	succCode := utilities.RF_CODE_200
	msg := utilities.GetMessageWithStatusCode(succCode)
	return notification.NewResponseWithModel(succCode, msg, nil)

}
