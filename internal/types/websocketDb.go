package types

import (
	"time"

	"gorm.io/gorm"
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

type Wcdma struct {
	CellID_Wcdma    int64 `json:"cellID"`
	Rsrp_Wcdma      int64 `db:"rsrp" json:"rsrp"`
	Rssi_Wcdma      int64 `db:"rssi" json:"rssi"`
	Rsrq_Wcdma      int64 `db:"rsrq" json:"rsrq"`
	Rssnr_Wcdma     int64 `db:"rssnr" json:"rssnr"`
	Cqi_Wcdma       int64 `db:"cqi" json:"cqi"`
	Bandwidth_Wcdma int64 `db:"bandwidth" json:"bandwidth"`
}

type Gsm struct {
	CellID_Gsm    int64 `json:"cellID"`
	Rsrp_Gsm      int64 `db:"rsrp" json:"rsrp"`
	Rssi_Gsm      int64 `db:"rssi" json:"rssi"`
	Rsrq_Gsm      int64 `db:"rsrq" json:"rsrq"`
	Rssnr_Gsm     int64 `db:"rssnr" json:"rssnr"`
	Cqi_Gsm       int64 `db:"cqi" json:"cqi"`
	Bandwidth_Gsm int64 `db:"bandwidth" json:"bandwidth"`
}

type Lte struct {
	CellID_Lte    int64 `json:"cellID"`
	Rsrp_Lte      int64 `db:"rsrp" json:"rsrp"`
	Rssi_Lte      int64 `db:"rssi" json:"rssi"`
	Rsrq_Lte      int64 `db:"rsrq" json:"rsrq"`
	Rssnr_Lte     int64 `db:"rssnr" json:"rssnr"`
	Cqi_Lte       int64 `db:"cqi" json:"cqi"`
	Bandwidth_Lte int64 `db:"bandwidth" json:"bandwidth"`
}

type Nr struct {
	CellID_Nr    int64 `json:"cellID"`
	Rsrp_Nr      int64 `db:"rsrp" json:"rsrp"`
	Rssi_Nr      int64 `db:"rssi" json:"rssi"`
	Rsrq_Nr      int64 `db:"rsrq" json:"rsrq"`
	Rssnr_Nr     int64 `db:"rssnr" json:"rssnr"`
	Cqi_Nr       int64 `db:"cqi" json:"cqi"`
	Bandwidth_Nr int64 `db:"bandwidth" json:"bandwidth"`
}

type MessageToData2 struct {
	gorm.Model
	UUID      string    `json:"UUID"`
	Time      time.Time `json:"time"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Wcdma      Wcdma      `json:"Wcdma" gorm:"embedded"`
	Gsm       Gsm       `json:"gsm" gorm:"embedded"`
	Lte       Lte       `json:"lte" gorm:"embedded"`
	Nr        Nr        `json:"nr" gorm:"embedded"`
}


!Оптимизация избыточности!

 *!ключ - global id, к нему линковать всю информацию!*
*/

type GsmData struct {
	ID            uint   `gorm:"primary_key"`
	RequestID     uint   `gorm:"index"`
	Registered    bool   `json:"registered"`
	Mcc           uint16 `json:"mcc"`
	Mnc           uint16 `json:"mnc"`
	Lac           int32  `json:"lac"`
	Cid           int32  `json:"cid"`
	Arfcn         uint16 `json:"arfcn"`
	Bsic          uint32 `json:"bsic"`
	Rssi          int32 `json:"rssi"`
	TimingAdvance uint32 `json:"timingAdvance"`
	BitErrorRate  uint32 `json:"bitErrorRate"`
}

type WcdmaData struct {
	ID           uint   `gorm:"primary_key"`
	RequestID    uint   `gorm:"index"`
	Registered   bool   `json:"registered"`
	Mcc          uint16 `json:"mcc"`
	Mnc          uint16 `json:"mnc"`
	Lac          int32  `json:"lac"`
	Cid          int32  `json:"cid"`
	Rssi         int32 `json:"rssi"`
	Psc          int32  `json:"psc"`
	Uarfcn       uint16 `json:"uarfcn"`
	Rscp         int32  `json:"rscp"`
	Ecno         int32  `json:"ecno"`
	Level        uint8  `json:"level"`
	BitErrorRate uint32 `json:"bitErrorRate"`
}

type LteData struct {
	ID            uint   `gorm:"primary_key"`
	RequestID     uint   `gorm:"index"`
	Type          string `json:"type"`
	Registered    bool   `json:"registered"`
	Mcc           uint16 `json:"mcc"`
	Mnc           uint16 `json:"mnc"`
	Ci            uint32 `json:"ci"`
	Pci           uint16 `json:"pci"`
	Tac           uint16 `json:"tac"`
	Earfcn        uint32 `json:"earfcn"`
	Bandwidth     uint32 `json:"bandwidth"`
	Rsrp          int16  `json:"rsrp"`
	Rssi          int16  `json:"rssi"`
	Rsrq          int32  `json:"rsrq"`
	Rssnr         int32  `json:"rssnr"`
	Cqi           int32  `json:"cqi"`
	TimingAdvance int32  `json:"timingAdvance"`
}

type NRData struct {
	ID         uint   `gorm:"primary_key"`
	RequestID  uint   `gorm:"index"`
	Type       string `json:"type"`
	Registered bool   `json:"registered"`
	Mcc        string `json:"mcc"`
	Mnc        string `json:"mnc"`
	Nci        int64  `json:"nci"`
	Nrarfcn    int32  `json:"nrarfcn"`
	Pci        int16  `json:"pci"`
	Tac        uint32 `json:"tac"`
	Bands      int    `json:"bands"`
}

type Message2 struct {
	gorm.Model
	JWT       string      `json:"jwt"`
	UUID      string      `json:"UUID"`
	Time      time.Time   `json:"time"`
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
	Operator  string      `json:"operator"`
	Wcdma     []WcdmaData `json:"wcdma" gorm:"foreignkey:RequestID"`
	Gsm       []GsmData   `json:"gsm" gorm:"foreignkey:RequestID"`
	Lte       []LteData   `json:"lte" gorm:"foreignkey:RequestID"`
	Nr        []NRData    `json:"nr" gorm:"foreignkey:RequestID"`
}
