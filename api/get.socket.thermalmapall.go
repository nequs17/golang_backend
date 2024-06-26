package api

import (
	"backend/internal/database"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func SocketThermalOut(w http.ResponseWriter, r *http.Request) {
	/*
		sessions, _ := cookie.Store.Get(r, "session-name")
		if sessions.Values["authenticated"] != true {
			net.Respond(w, http.StatusBadRequest, net.Msg{
				"error": "You are not auth",
			})
			return
		} else {
	*/

	uuid := "TEST"

	var data []Data

	if err := database.DB.Table("message_to_data").Where("uuid = ?", uuid).Find(&data).Error; err != nil {
		http.Error(w, fmt.Sprintf("Failed to query data: %v", err), http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		http.Error(w, "No data found for the given ID", http.StatusNotFound)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(data); err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
	// }

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/json")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		if err := json.NewEncoder(gz).Encode(data); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
			return
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
		}
	}
}
