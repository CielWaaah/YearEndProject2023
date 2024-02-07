package model

import "github.com/jinzhu/gorm"

type Dormitory struct {
	gorm.Model
	AmMeter_ID string `json:"am_meter_id" form:"am_meter_id"`
}
