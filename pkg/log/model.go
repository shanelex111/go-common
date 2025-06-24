package log

import (
	"encoding/json"
	"time"

	"maps"

	"github.com/sirupsen/logrus"
)

type logger struct {
	Level string    `json:"level"`
	Time  time.Time `json:"time"`
}

type customJSONFormatter struct {
}

func (f *customJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data))
	maps.Copy(data, entry.Data)

	logMeta := &logger{
		Level: entry.Level.String(),
		Time:  entry.Time,
	}

	data["log"] = logMeta

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return append(serialized, '\n'), nil
}
