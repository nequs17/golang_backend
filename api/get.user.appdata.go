package api

import (
	"backend/cookie"
	"backend/internal/database"
	"backend/internal/types"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type MessageApp struct {
	gorm.Model
	IDMessageApp uint             `gorm:"index;unique"`
	JWT          string           `json:"jwt"`
	UUID         string           `json:"uuid"`
	TrafficData  []TrafficDataApp `json:"trafficData" gorm:"foreignKey:IDMessageApp"`
}

type TrafficDataApp struct {
	gorm.Model
	AppName     string `json:"appName"`
	TotalBytes  int64  `json:"totalBytes"`
	MobileBytes int64  `json:"mobileBytes"`
	WifiBytes   int64  `json:"wifiBytes"`
	RxBytes     int64  `json:"rxBytes"`
	TxBytes     int64  `json:"txBytes"`
}

// GetAppTrafic godoc
// @Summary Get App Trafic
// @Tags user
// @ID GetAppTrafic
// @Accept json
// @Produce json
// @Success 200 Traffic data
// @Failure 400 {object} string "No data"
// @Failure 500 {string} string "Failed to encode data to JSON: [Error Message]"
// @Router /api/user/getapptrafic [get]
func GetAppTraffic(w http.ResponseWriter, r *http.Request) {

	sessions, _ := cookie.Store.Get(r, "session-name")
	authenticated, authOk := sessions.Values["authenticated"].(bool)
	uuid := sessions.Values["username"]
	if authOk && authenticated {
		var appUserData []types.Request

		if err := database.DB.Preload("requests").Where("uuid = ?", uuid).Find(&appUserData).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "No data", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(appUserData); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode data to JSON: %v", err), http.StatusInternalServerError)
		}
	}
}

/*

db.Preload("TrafficData").Find(&requests)
  for _, req := range requests {
    fmt.Printf("Request ID: %d, JWT: %s, UUID: %s\n", req.ID, req.JWT, req.UUID)
    for _, traffic := range req.TrafficData {
      fmt.Printf("\tApp Name: %s, Total Bytes: %d, Mobile Bytes: %d, Wifi Bytes: %d, Rx Bytes: %d, Tx Bytes: %d\n",
        traffic.AppName, traffic.TotalBytes, traffic.MobileBytes, traffic.WifiBytes, traffic.RxBytes, traffic.TxBytes)
    }
  }


*/
