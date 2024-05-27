package api

import (
	"backend/internal/database"
	"backend/internal/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	//log "github.com/sirupsen/logrus"
)

/*
Time      time.Time `json:"time"`

	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	Rsrp      string    `json:"rsrp"`
	Rssi      string    `json:"rssi"`
	Rsrq      string    `json:"rsrq"`
	Rssnr     string    `json:"rssnr"`
	Cqi       string    `json:"cqi"`
	Bandwidth string    `json:"bandwidth"`
	CellID    string    `json:"cellID"`
*/
type Data struct {
	UUID      string    `db:"uuid"`
	Time      time.Time `db:"time" json:"time"`
	Latitude  string    `db:"latitude" json:"latitude"`
	Longitude string    `db:"longitude" json:"longitude"`
	Rsrp      string    `db:"rsrp" json:"rsrp"`
	Rssi      string    `db:"rssi" json:"rssi"`
	Rsrq      string    `db:"rsrq" json:"rsrq"`
	Rssnr     string    `db:"rssnr" json:"rssnr"`
	Cqi       string    `db:"cqi" json:"cqi"`
	Bandwidth string    `db:"bandwidth" json:"bandwidth"`
	CellID    string    `db:"cell_id" json:"cellID"`
}

func SocketThermalOut(w http.ResponseWriter, r *http.Request) {

	uuid := "TEST"

	user := types.Account{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println(err)
		return
	}

	var data []Data

	if err := database.DB.Table("message_to_data").Where("uuid = ?", uuid).Find(&data).Error; err != nil {
		http.Error(w, fmt.Sprintf("Failed to query data: %v", err), http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		http.Error(w, "No data found for the given ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
	}
}
