package redis

type RedisHook struct{}

type logEntry struct {
	StartAt int64    `json:"start_at"`
	Latency int64    `json:"latency"`
	Cmds    []string `json:"cmds"`
	EndAt   int64    `json:"end_at"`
	Msg     string   `json:"msg"`
}
