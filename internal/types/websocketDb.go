package types

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MessageToData struct {
	gorm.Model
	UUID      string    `json:"UUID"`
	Time      time.Time `json:"time"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Rsrp      int64     `json:"rsrp"`
	Rssi      int64     `json:"rssi"`
	Rsrq      int64     `json:"rsrq"`
	Rssnr     int64     `json:"rssnr"`
	Cqi       int64     `json:"cqi"`
	Bandwidth int64     `json:"bandwidth"`
	CellID    int64     `json:"cellID"`
}
