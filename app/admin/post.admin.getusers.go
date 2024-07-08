package appAdmin

import (
	"backend/cookie"
	"backend/internal/database"
	"backend/internal/net"
	"encoding/json"
	"fmt"
	"net/http"
)

type Email string

type Data_users struct {
	Email    Email  `json:"email" gorm:"primaryKey"`
	UUID     string `json:UUID gorm:"primaryKey"`
	Group    string `json:group`
	Password string `json:"password"`
	Verifed  bool   `json:"verifed"`
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")
	if sessions.Values["role"].(int) < 10 && sessions.Values["authenticated"] == false {
		net.Respond(w, http.StatusBadRequest, net.Msg{
			"error": "You not admin",
		})
		return
	} else {

		var data []Data_users

		if err := database.DB.Table("message_to_data").Find(&data).Error; err != nil {
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
}
