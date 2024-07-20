package api

import (
	"backend/cookie"
	"backend/internal/database"
	"backend/internal/net"
	"backend/internal/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SocketThermalOut(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")
	if sessions.Values["authenticated"] != true {
		net.Respond(w, http.StatusBadRequest, net.Msg{
			"error": "You are not auth",
		})
		return
	} else {
		var requests []types.Message2
		if err := database.DB.Preload("Cdma").Preload("Gsm").Preload("Lte").Preload("Nr").Find(&requests).Error; err != nil {
			log.Fatalf("failed to preload requests: %v", err)
		}

		if len(requests) == 0 {
			http.Error(w, "No data found for the given ID", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(requests); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
		}
	}
}
