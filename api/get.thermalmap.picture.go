package api

import (
	"backend/internal/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.bug.st/serial.v1"
	//log "github.com/sirupsen/logrus"
)

// ThermalMapPicture retrieves thermal data based on query parameters.
// @Summary Retrieves thermal data based on query parameters
// @Description This endpoint retrieves thermal data based on query parameters and returns an image.
// @ID ThermalMapPicture
// @Accept json
// @Produce png
// @Param start_time query string false "Start Time in RFC3339 format (example: 2024-05-18T18:15:00.000+03:00)"
// @Param end_time query string false "End Time in RFC3339 format (example: 2024-05-18T18:16:00.000+03:00)"
// @Success 200 {file} file "Image representing the thermal data"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "No data found"
// @Failure 500 {string} string "Failed to query or encode data"
// @Router /api/picture/thermalmappic [get]
func ThermalMapPicture(w http.ResponseWriter, r *http.Request) {
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

	if err := db.Find(&data).Error; err != nil {
		http.Error(w, fmt.Sprintf("Failed to query data: %v", err), http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}
	jsondata, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error in json creating: %v", err)
	}
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open("/dev/picture", mode)

	if err != nil {
		log.Fatalf("Error in open port: %v", err)
	}

	defer port.Close()

	_, err = port.Write(jsondata)
	if err != nil {
		http.Error(w, "Ошибка при передачи данных", http.StatusInternalServerError)
	}

	buff := make([]byte, 1024)
	n, err := port.Read(buff)
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "image/png")
	w.Write(buff[:n])
}
