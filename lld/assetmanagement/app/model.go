package app

import "time"

type Asset struct {
	AssetID      int       `json:"asset_id"`
	TimeStamp    time.Time `json:"timestamp"`
	BatteryLevel int       `json:"battery_level"`
}
