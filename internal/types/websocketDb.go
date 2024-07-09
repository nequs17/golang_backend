package types

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MessageToData struct {
	gorm.Model
	UUID       string    `json:"UUID"`
	Time       time.Time `json:"time"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Rsrp       int64     `json:"rsrp"`
	Rssi       int64     `json:"rssi"`
	Rsrq       int64     `json:"rsrq"`
	Rssnr      int64     `json:"rssnr"`
	Cqi        int64     `json:"cqi"`
	Bandwidth  int64     `json:"bandwidth"`
	CellID     int64     `json:"cellID"`
	PhysCellID int64     `json:"PhysCellID"`
}

/*

package types

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Cdma struct {
	CellID    int64   `json:"cellID"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Rsrp      int64   `db:"rsrp" json:"rsrp"`
	Rssi      int64   `db:"rssi" json:"rssi"`
	Rsrq      int64   `db:"rsrq" json:"rsrq"`
	Rssnr     int64   `db:"rssnr" json:"rssnr"`
	Cqi       int64   `db:"cqi" json:"cqi"`
	Bandwidth int64   `db:"bandwidth" json:"bandwidth"`
}

type Gsm struct {
	CellID    int64   `json:"cellID"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Rsrp      int64   `db:"rsrp" json:"rsrp"`
	Rssi      int64   `db:"rssi" json:"rssi"`
	Rsrq      int64   `db:"rsrq" json:"rsrq"`
	Rssnr     int64   `db:"rssnr" json:"rssnr"`
	Cqi       int64   `db:"cqi" json:"cqi"`
	Bandwidth int64   `db:"bandwidth" json:"bandwidth"`
}

type Lte struct {
	CellID    int64   `json:"cellID"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Rsrp      int64   `db:"rsrp" json:"rsrp"`
	Rssi      int64   `db:"rssi" json:"rssi"`
	Rsrq      int64   `db:"rsrq" json:"rsrq"`
	Rssnr     int64   `db:"rssnr" json:"rssnr"`
	Cqi       int64   `db:"cqi" json:"cqi"`
	Bandwidth int64   `db:"bandwidth" json:"bandwidth"`
}

type Nr struct {
	CellID    int64   `json:"cellID"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Rsrp      int64   `db:"rsrp" json:"rsrp"`
	Rssi      int64   `db:"rssi" json:"rssi"`
	Rsrq      int64   `db:"rsrq" json:"rsrq"`
	Rssnr     int64   `db:"rssnr" json:"rssnr"`
	Cqi       int64   `db:"cqi" json:"cqi"`
	Bandwidth int64   `db:"bandwidth" json:"bandwidth"`
}

type MessageToData2 struct {
	gorm.Model
	UUID      string    `json:"UUID"`
	Time      time.Time `json:"time"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Cdma      Cdma      `json:"cdma"`
	Gsm       Gsm       `json:"gsm"`
	Lte       Lte       `json:"lte"`
	Nr        Nr        `json:"nr"`
}

!Оптимизация избыточности!

 *!ключ - global id, к нему линковать всю информацию!*

*/
