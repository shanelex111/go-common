package request

import "github.com/shanelex111/go-common/third_party/geo"

type logEntryRequest struct {
	Method    string              `json:"method"`
	Path      string              `json:"path"`
	Header    map[string][]string `json:"header"`
	Body      string              `json:"body"`
	Timestamp int64               `json:"timestamp"`
}

type logEntryResponse struct {
	Status    int                 `json:"status"`
	Latency   int64               `json:"latency"`
	Header    map[string][]string `json:"header"`
	Body      string              `json:"body"`
	Timestamp int64               `json:"timestamp"`
}

type TokenInfo struct {
	Account *TokenInfoAccount `json:"account"`
	Access  *TokenInfoAccess  `json:"access"`
	Device  *TokenInfoDevice  `json:"device"`
	Geo     *geo.GeoCity      `json:"geo"`
}

type TokenInfoAccount struct {
	ID uint `json:"id"`
}
type TokenInfoAccess struct {
	Token            string `json:"token"`
	ExpiredAt        int64  `json:"expired_at"`
	Refresh          string `json:"refresh"`
	RefreshExpiredAt int64  `json:"refresh_expired_at"`
}

type TokenInfoDevice struct {
	DeviceID    string `json:"device_id"`
	DeviceModel string `json:"device_model"`
	DeviceType  string `json:"device_type"`
	AppVersion  int    `json:"app_version"`
	CreatedAt   int64  `json:"created_at"`
}
