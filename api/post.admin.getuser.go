package api

import (
	"backend/cookie"
	"backend/internal/database"
	"backend/internal/net"
	"encoding/json"
	"fmt"
	"net/http"
)

type Data_users struct {
	Email   string `json:"email"`
	UUID    string `json:UUID`
	Group   string `json:group`
	Verifed bool   `json:"verifed"`
}

// AllUsers godoc
// @Summary Get all users
// @Description Get all users from the database
// @Tags users
// @Produce  json
// @Success 200 {array} Data_users
// @Failure 400 {object} string "You not admin"
// @Failure 404 {string} string "No data found"
// @Failure 500 {string} string "Failed to query data: [Error Message]"
// @Router /api/admin/users [post]
func AllUsers(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")
	fmt.Println(sessions.Values["role"])
	if sessions.Values["role"] != "Admin" {
		net.Respond(w, http.StatusBadRequest, net.Msg{
			"error": "You not admin",
		})
		return
	}
	if sessions.Values["role"] == "Admin" && sessions.Values["authenticated"] == true {

		var data []Data_users

		// Query the database for all users
		if err := database.DB.Table("accounts").Find(&data).Error; err != nil {
			http.Error(w, fmt.Sprintf("Failed to query data: %v", err), http.StatusInternalServerError)
			return
		}

		// Check if no data is found
		if len(data) == 0 {
			http.Error(w, "No data found", http.StatusNotFound)
			return
		}

		// Set response header to JSON and encode data
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
		}
	}
}
