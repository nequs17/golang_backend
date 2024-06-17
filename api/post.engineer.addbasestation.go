package api

import (
	"backend/cookie"
	"backend/internal/net"
	"backend/internal/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// AddBaseStation godoc
// @Summary Change User Role
// @Tags engineer
// @ID AddBaseStation
// @Accept json
// @Produce json
// @Param Latitude query string false "lat"
// @Param Longitude query string false "lon"
// @Param CellID query string false "cellid"
// @Success 200 Role change
// @Failure 400 {object} string "You not engineer or admin"
// @Failure 404 {string} string "No data found"
// @Failure 500 {string} string "Failed to query data: [Error Message]"
// @Router /api/engineer/addbasestation [post]
func AddBaseStation(w http.ResponseWriter, r *http.Request) {
	sessions, _ := cookie.Store.Get(r, "session-name")
	authenticated, authOk := sessions.Values["authenticated"].(bool)
	role, roleOk := sessions.Values["role"].(int)

	if authOk && roleOk && role >= 5 && authenticated {

		lat := r.URL.Query().Get("lat")
		lon := r.URL.Query().Get("lon")
		cellID := r.URL.Query().Get("cellid")

		var station types.BaseStation

		latFloat, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			fmt.Errorf("Error in parse float (addbasestation - lat): %v", err)
		}
		lonFloat, err := strconv.ParseFloat(lon, 64)
		if err != nil {
			fmt.Errorf("Error in parse float (addbasestation - lon): %v", err)
		}
		cellIDint, err := strconv.ParseInt(cellID, 64, 64)
		if err != nil {
			fmt.Errorf("Error in parse int (addbasestation - cellID): %v", err)
		}
		if !station.Addbasestation(latFloat, lonFloat, cellIDint) {
			net.Respond(w, http.StatusBadRequest, net.Msg{
				"error": "error adding BS",
			})
			return
		} else {
			json.NewEncoder(w).Encode(map[string]string{"message": "BS has been successfully added"})
		}

	} else {
		net.Respond(w, http.StatusBadRequest, net.Msg{
			"error": "You are not authorized or not an admin",
		})
		return
	}
}
