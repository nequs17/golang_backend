package api

import (
	"backend/cookie"
	"backend/internal/database"
	"backend/internal/net"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

// ChangeRole godoc
// @Summary Change User Role
// @Tags admin
// @ID ChangeRole
// @Accept json
// @Produce json
// @Param email query string false "email"
// @Param role query int false "role"
// @Success 200 Role change
// @Failure 400 {object} string "You not admin"
// @Failure 404 {string} string "No data found"
// @Failure 500 {string} string "Failed to query data: [Error Message]"
// @Router /api/admin/changerole [get]
func ChangeRole(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")

	// Проверяем, аутентифицирован ли пользователь и установлено ли значение роли в сессии
	authenticated, authOk := sessions.Values["authenticated"].(bool)
	role, roleOk := sessions.Values["role"].(int)

	if authOk && roleOk && role >= 10 && authenticated {

		email := r.URL.Query().Get("email")
		rolestr := r.URL.Query().Get("role")

		role, err := strconv.Atoi(rolestr)
		if err != nil {
			http.Error(w, "Invalid role parameter", http.StatusBadRequest)
			return
		}

		var editUser Data_users

		if err := database.DB.Table("accounts").Where("email = ?", email).First(&editUser).Update("group", role).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "user not exists", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "Role update successfully"})
	} else {
		net.Respond(w, http.StatusBadRequest, net.Msg{
			"error": "You are not authorized or not an admin",
		})
		return
	}
}
