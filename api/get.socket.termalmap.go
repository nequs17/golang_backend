package api

import (
	"backend/internal/types"
	"net"
	"os"
	"sync"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tcaine/twamp"
	//log "github.com/sirupsen/logrus"
)

type Message struct {
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

var connPool = sync.Pool{
	New: func() interface{} {
		conn, _, err := websocket.DefaultDialer.Dial("ws://"+os.Getenv("SERVER_HOST")+":"+os.Getenv("SERVER_PORT")+"/api/sockets/termalmap", nil)
		if err != nil {
			log.Fatalf("failed to dial websocket server: %v", err)
		}
		return conn
	},
}

func SocketThermal(w http.ResponseWriter, r *http.Request) {
	twamp.NewClient()
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Ip : ", ip)
	fmt.Println("Port:", port)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	user := &types.Account{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println(err)
		return
	}

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
		fmt.Println("Received message:", msg)
		/*DataToDB := &types.MessageToData{
			UUID: user.UUID,
			//UUID:      "TEST",
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
		database.DB.Create(&DataToDB)*/
	}
}
