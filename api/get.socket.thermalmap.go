package api

import (
	"backend/internal/database"
	"backend/internal/net"
	"backend/internal/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	//log "github.com/sirupsen/logrus"
)

type Message struct {
	JWT       string    `json:"jwt"`
	UUID      string    `json:"uuid"`
	Time      time.Time `json:"time"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Rsrp      int64     `json:"rsrp"`
	Rssi      int64     `json:"rssi"`
	Rsrq      int64     `json:"rsrq"`
	Rssnr     int64     `json:"rssnr"`
	Cqi       int64     `json:"cqi"`
	Bandwidth int64     `json:"bandwidth"`
	CellID    int64     `json:"cellID"`
}

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
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(err)
			continue
		}

		token := types.Token{JWT: msg.JWT}

		if result, _ := token.Verify(); !result {
			net.Respond(w, http.StatusForbidden, net.Msg{
				"error": "Unauthorized access blocked",
			})
			return
		}

		fmt.Println("Received message:", msg)
		DataToDB := &types.MessageToData{
			UUID:      msg.UUID,
			Time:      msg.Time,
			Latitude:  msg.Latitude,
			Longitude: msg.Longitude,
			Rsrp:      msg.Rsrp,
			Rssi:      msg.Rssi,
			Rsrq:      msg.Rsrq,
			Rssnr:     msg.Rssnr,
			Cqi:       msg.Cqi,
			Bandwidth: msg.Bandwidth,
			CellID:    msg.CellID,
		}
		database.DB.Create(&DataToDB)
		conn.SetReadDeadline(time.Now().Add(300 * time.Second))
	}
}
