package types

type TrafficData struct {
	ID          uint   `gorm:"primary_key"`
	RequestID   uint   `gorm:"index"`
	AppName     string `gorm:"type:varchar(100)"`
	TotalBytes  uint64
	MobileBytes uint64
	WifiBytes   uint64
	RxBytes     uint64
	TxBytes     uint64
}

type Request struct {
	ID              uint          `gorm:"primary_key"`
	JWT             string        `gorm:"type:varchar(255)"`
	UUID            string        `gorm:"type:varchar(255)"`
	Role            string        `gorm:"type:varchar(255)"`
	TrafficDataTEST []TrafficData `gorm:"foreignkey:RequestID"`
}
