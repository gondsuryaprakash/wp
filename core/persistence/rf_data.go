package persistence

import (
	"waetherproject.com/core/domain"
	"waetherproject.com/model"
)

type RfPersistence interface {
	// GetRFByDateRange()
	GetRFById(id int) (v *domain.RfDataDomain, err error)
	GetRFByLatAndLong(lat, long string) (v *domain.RfDataDomain, err error)
	AddRF(m *domain.RfDataDomain) error
	UpdateRFById(m *domain.RfDataDomain, cols ...string) (err error)
	DeleteRF(id int) error
}

func NewRfPersistence() RfPersistence {
	return &model.BeegoRFData{}
}
