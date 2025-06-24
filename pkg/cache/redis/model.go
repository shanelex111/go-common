package redis

type RedisHook struct{}

type logEntry struct {
	StartAt int64    `json:"start_at"`
	Elapsed int64    `json:"elapsed"`
	Cmds    []string `json:"cmds"`
	EndAt   int64    `json:"end_at"`
}
