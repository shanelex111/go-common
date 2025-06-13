package request

type logEntryRequest struct {
	Method string              `json:"method"`
	Path   string              `json:"path"`
	Header map[string][]string `json:"header"`
	Body   string              `json:"body"`
}

type logEntryResponse struct {
	Status  int                 `json:"status"`
	Latency int64               `json:"latency"`
	Header  map[string][]string `json:"header"`
	Body    string              `json:"body"`
}
