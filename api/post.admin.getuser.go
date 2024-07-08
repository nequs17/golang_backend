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
// @Summary Get all user
// @Description Get all users from the database
// @Tags admin
// @Produce  json
// @Success 200 {array} Data_users
// @Failure 400 {object} string "You not admin"
// @Failure 404 {string} string "No data found"
// @Failure 500 {string} string "Failed to query data: [Error Message]"
// @Router /api/admin/users [post]
func AllUsers(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")

	// Проверяем, аутентифицирован ли пользователь и установлено ли значение роли в сессии
	authenticated, authOk := sessions.Values["authenticated"].(bool)
	role, roleOk := sessions.Values["role"].(int)

	if authOk && roleOk && role >= 10 && authenticated {
		var data []Data_users

		// Запрос данных из базы данных для всех пользователей
		if err := database.DB.Table("accounts").Find(&data).Error; err != nil {
			http.Error(w, fmt.Sprintf("Failed to query data: %v", err), http.StatusInternalServerError)
			return
		}

		// Проверка наличия данных
		if len(data) == 0 {
			http.Error(w, "No data found", http.StatusNotFound)
			return
		}

		// Устанавливаем заголовок ответа на JSON и кодируем данные
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
		}
	} else {
		net.Respond(w, http.StatusBadRequest, net.Msg{
			"error": "You are not authorized or not an admin",
		})
		return
	}
}

/*

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

*/
