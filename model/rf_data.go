package model

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"waetherproject.com/core/domain"
	"waetherproject.com/util/logger"
)

type RfData struct {
	Id        int       `orm:"column(id);auto" json:"id"`
	Latitude  *string   `orm:"column(latitude)" json:"latitude"`
	Longitude *string   `orm:"column(longitude)" json:"longitude"`
	RF        *string   `orm:"column(rf)" json:"rf"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}

type BeegoRFData struct {
}

func init() {
	funcName := "model.init"
	logger.I(funcName)
	orm.RegisterModel(new(RfData))
}

func (rf *RfData) TableName() string {
	return "rf_data"
}

// AddRF with RF struct
func (c *BeegoRFData) AddRF(m *domain.RfDataDomain) error {
	funcName := "model.AddRF"
	orm := orm.NewOrm()
	ormdata := c.ConvertDomainToModel(m)
	if _, err := orm.Insert(ormdata); err != nil {
		logger.E(funcName, err)
		return err
	}
	return nil
}

//GetRFByLatAndLong fetch the RF value by Lat and Long
func (c *BeegoRFData) GetRFByLatAndLong(lat, long string) (v *domain.RfDataDomain, err error) {
	funcName := "model.GetRFByLatAndLong"
	logger.I(funcName)
	orm := orm.NewOrm()
	rfData := new(RfData)
	m := &RfData{}
	if err = orm.QueryTable(rfData).Filter("latitude", lat).Filter("longitude", long).One(m); err != nil {
		logger.E(funcName, err)
		return nil, err
	}
	v = c.ConvertModelToDomain(m)
	return
}

//GetRFByDateRange fetch the RF value by the Date Range
func (c *BeegoRFData) GetRFByDateRange() {

}

// GetRFById
func (c *BeegoRFData) GetRFById(id int) (v *domain.RfDataDomain, err error) {

	funcName := "model.GetRFById"
	logger.I(funcName)
	orm := orm.NewOrm()

	m := &RfData{Id: id}
	if err = orm.Read(m); err != nil {
		logger.E(funcName, err)
		return nil, err
	}

	v = c.ConvertModelToDomain(m)
	return v, nil

}

// Update RF By Id
func (c *BeegoRFData) UpdateRFById(m *domain.RfDataDomain, cols ...string) (err error) {
	funcName := "model.UpdateRFById"
	logger.I(funcName)
	orm := orm.NewOrm()
	v := &RfData{Id: m.Id}
	if err := orm.Read(v); err == nil {
		var num int64
		l := c.ConvertDomainToModel(m)
		if num, err = orm.Update(l, cols...); err != nil {
			logger.E(funcName, err, num)
		}
		logger.I(funcName, num)
		return nil
	}
	return
}

func (c *BeegoRFData) DeleteRF(id int) error {
	funcName := "model.DeleteRF"
	logger.I(funcName)
	o := orm.NewOrm()

	m := &RfData{Id: id}
	if err := o.Read(m); err == nil {
		if _, err1 := o.Delete(m); err1 != nil {
			logger.E(funcName, err1)
			return err1
		}

	} else {
		logger.E(funcName, err)
		return err
	}

	return nil

}

func (c *BeegoRFData) ConvertModelToDomain(v *RfData) *domain.RfDataDomain {
	return &domain.RfDataDomain{
		Id:        v.Id,
		Latitude:  v.Latitude,
		Longitude: v.Longitude,
		RF:        v.RF,
		Created:   v.Created,
		Updated:   v.Updated,
	}

}

func (c *BeegoRFData) ConvertDomainToModel(v *domain.RfDataDomain) *RfData {
	return &RfData{
		Id:        v.Id,
		Latitude:  v.Latitude,
		Longitude: v.Longitude,
		RF:        v.RF,
		Created:   v.Created,
		Updated:   v.Updated,
	}
}
