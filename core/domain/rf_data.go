package domain

import "time"

type RfDataDomain struct {
	Id        int
	Latitude  *string
	Longitude *string
	RF        *string
	Created   time.Time
	Updated   time.Time
}

type RfDataDomainReq struct {
	Latitude  *string `json:"latitude"`
	Longitude *string `json:"longitude"`
	RF        *string `json:"rf"`
}

type RfDataDomainReqForUpdate struct {
	Id        int     `json:"id"`
	Latitude  *string `json:"latitude"`
	Longitude *string `json:"longitude"`
	RF        *string `json:"rf"`
}
