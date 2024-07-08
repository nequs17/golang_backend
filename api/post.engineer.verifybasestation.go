package api

import (
	"backend/cookie"
	"backend/internal/database"
	"backend/internal/net"
	"backend/internal/types"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

// verifybasestation godoc
// @Summary Change User Role
// @Tags engineer
// @ID verifybasestation
// @Accept json
// @Produce json
// @Param Latitude query string false "lat"
// @Param Longitude query string false "lon"
// @Param CellID query string false "cellid"
// @Success 200 Role change
// @Failure 400 {object} string "You not engineer or admin"
// @Failure 404 {string} string "No data found"
// @Failure 500 {string} string "Failed to query data: [Error Message]"
// @Router /api/engineer/verifybasestation [post]
func VerifyBaseStation(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")
	authenticated, authOk := sessions.Values["authenticated"].(bool)
	role, roleOk := sessions.Values["role"].(int)

	if authOk && roleOk && role >= 5 && authenticated {

		cellID := r.URL.Query().Get("cellid")

		cellIDint, err := strconv.ParseInt(cellID, 64, 64)

		var bs_temp types.BaseStation

		if err != nil {
			fmt.Errorf("Error in parse int (addbasestation - cellID): %v", err)
		} else {
			json.NewEncoder(w).Encode(map[string]string{"message": "BS has been successfully added"})
		}
		if err := database.DB.Table("base_station").Where("cell_id = ?", cellIDint).First(&bs_temp).Update("verify", true).Error; errors.Is(err, gorm.ErrRecordNotFound) {
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
