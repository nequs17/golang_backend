package types

import (
	"backend/internal/database"
	"time"

	"github.com/jinzhu/gorm"
)

type BaseStation struct {
	gorm.Model
	Time      time.Time `json:"time"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CellID    int64     `json:"cellID"`
	Verify    bool      `json:"Verify"`
}

/*
шина полосы - может поменятся
номер pci - может поменятся
проверка БС на актуальность информации
перезаписанные данные требуется хранить
запись нескольких технологий (gsm, let, nr и т.п)
триггерные изменения информации

фиксировать глобальный cellid, если параметр изменился -> запись в срез данных и времени (проверка по таймеру)
*/
func (station *BaseStation) Veryfibasestation(lat float64, lon float64, CellID int64) bool {
	notExists := database.DB.Table("accounts").Where("Latitude = ?", station.Latitude).Where("Longitude = ?", station.Longitude).Where("CellID = ?", station.CellID).First(&Account{}).RecordNotFound()
	if !notExists {
		return false
	} else {
	}
	return false
}

func (station *BaseStation) Addbasestation(lat float64, lon float64, CellID int64) bool {
	notExists := database.DB.Table("accounts").Where("Latitude = ?", station.Latitude).Where("Longitude = ?", station.Longitude).Where("CellID = ?", station.CellID).First(&Account{}).RecordNotFound()
	if !notExists {
		return false
	} else {
		station := BaseStation{
			Time:      time.Now(),
			Latitude:  lat,
			Longitude: lon,
			CellID:    CellID,
			Verify:    false,
		}
		database.DB.Create(&station)
		return false
	}
}
