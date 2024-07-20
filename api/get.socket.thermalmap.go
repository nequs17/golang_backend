package api

import (
	"backend/internal/database"
	"backend/internal/types"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	//log "github.com/sirupsen/logrus"
)

// ////////////////////////
/* old struct
type Message struct {
	JWT        string    `json:"jwt"`
	UUID       string    `json:"uuid"`
	Time       time.Time `json:"time"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	CellID     int64     `json:"cellID"`
	PhysCellID int64     `json:"PhysCellID"`
	Rsrp       int64     `json:"rsrp"`
	Rssi       int64     `json:"rssi"`
	Rsrq       int64     `json:"rsrq"`
	Rssnr      int64     `json:"rssnr"`
	Cqi        int64     `json:"cqi"`
	Bandwidth  int64     `json:"bandwidth"`
}
*/
var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 10 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SocketThermal(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		var msg types.Message2
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}
		token := types.Token{JWT: msg.JWT}
		isValid, err := token.Verify()

		if !isValid && err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Unauthorized: invalid token"))
		} else {
			/*
				fmt.Println("PhysCellId:", msg.PhysCellID)
				DataToDB := &types.MessageToData{
					UUID:       msg.UUID,
					Time:       msg.Time,
					Latitude:   msg.Latitude,
					Longitude:  msg.Longitude,
					Rsrp:       msg.Rsrp,
					Rssi:       msg.Rssi,
					Rsrq:       msg.Rsrq,
					Rssnr:      msg.Rssnr,
					Cqi:        msg.Cqi,
					Bandwidth:  msg.Bandwidth,
					CellID:     msg.CellID,
					PhysCellID: msg.PhysCellID,

						Cdma:      types.Cdma(msg.Cdma),
						Gsm:       types.Gsm(msg.Gsm),
						Lte:       types.Lte(msg.Lte),
						Nr:        types.Nr(msg.Nr),
				}
				database.DB.Create(&DataToDB)
				conn.SetReadDeadline(time.Now().Add(300 * time.Second))
			*/
			if err := database.DB.Create(&msg).Error; err != nil {
				log.Fatalf("failed to create request: %v", err)
			}
		}
	}
}
