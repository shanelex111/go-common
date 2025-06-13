package request

import "time"

type logEntryLogger struct {
	Level string    `json:"level"`
	Time  time.Time `json:"time"`
}
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

type logEntry struct {
	Log      *logEntryLogger   `json:"log"`
	Request  *logEntryRequest  `json:"request"`
	Response *logEntryResponse `json:"response"`
}
