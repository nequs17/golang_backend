package api

import (
	"backend/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	//log "github.com/sirupsen/logrus"
)

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

// SocketThermalOut retrieves thermal data based on query parameters.
// @Summary Retrieves thermal data based on query parameters
// @Description This endpoint retrieves thermal data based on query parameters.
// @ID socketThermalOut
// @Accept json
// @Produce json
// @Param uuid query string false "UUID"
// @Param start_time query string false "Start Time in RFC3339 format (example: 2024-05-18T18:15:00.000+03:00)"
// @Param end_time query string false "End Time in RFC3339 format (example: 2024-05-18T18:16:00.000+03:00)"
// @Param latitude query string false "Latitude"
// @Param longitude query string false "Longitude"
// @Param rsrp query string false "RSRP"
// @Param rssi query string false "RSSI"
// @Param rsrq query string false "RSRQ"
// @Param rssnr query string false "RSSNR"
// @Param cqi query string false "CQI"
// @Param bandwidth query string false "Bandwidth"
// @Param cell_id query string false "Cell ID"
// @Success 200 {array} Data "List of thermal data"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "No data found"
// @Failure 500 {string} string "Failed to query or encode data"
// @Router /api/sockets/thermalmapdata [get]
func SocketThermalOutByParams(w http.ResponseWriter, r *http.Request) {
	/*
		sessions, _ := cookie.Store.Get(r, "session-name")
		if sessions.Values["authenticated"] != true {
			http.Error(w, "You are not authenticated", http.StatusBadRequest)
			return
		}
	*/
	var data []Data
	db := database.DB.Table("message_to_data")

	query := r.URL.Query()

	if uuid := query.Get("uuid"); uuid != "" {
		db = db.Where("uuid = ?", uuid)
	}
	if startTimeStr := query.Get("start_time"); startTimeStr != "" {
		if startTime, err := time.Parse(time.RFC3339, startTimeStr); err == nil {
			db = db.Where("time >= ?", startTime)
		}
	}
	if endTimeStr := query.Get("end_time"); endTimeStr != "" {
		if endTime, err := time.Parse(time.RFC3339, endTimeStr); err == nil {
			db = db.Where("time <= ?", endTime)
		}
	}
	if latitude := query.Get("latitude"); latitude != "" {
		db = db.Where("latitude = ?", latitude)
	}
	if longitude := query.Get("longitude"); longitude != "" {
		db = db.Where("longitude = ?", longitude)
	}
	if rsrp := query.Get("rsrp"); rsrp != "" {
		db = db.Where("rsrp = ?", rsrp)
	}
	if rssi := query.Get("rssi"); rssi != "" {
		db = db.Where("rssi = ?", rssi)
	}
	if rsrq := query.Get("rsrq"); rsrq != "" {
		db = db.Where("rsrq = ?", rsrq)
	}
	if rssnr := query.Get("rssnr"); rssnr != "" {
		db = db.Where("rssnr = ?", rssnr)
	}
	if cqi := query.Get("cqi"); cqi != "" {
		db = db.Where("cqi = ?", cqi)
	}
	if bandwidth := query.Get("bandwidth"); bandwidth != "" {
		db = db.Where("bandwidth = ?", bandwidth)
	}
	if cellID := query.Get("cell_id"); cellID != "" {
		db = db.Where("cell_id = ?", cellID)
	}

	if err := db.Find(&data).Error; err != nil {
		http.Error(w, fmt.Sprintf("Failed to query data: %v", err), http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
	}
}

/*
	sessions, _ := cookie.Store.Get(r, "session-name")
	if sessions.Values["authenticated"] != true {
		net.Respond(w, http.StatusBadRequest, net.Msg{
			"error": "You are not auth",
		})
		return
	} else {
		uuid := "TEST"
		query := r.URL.Query()
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

*/
