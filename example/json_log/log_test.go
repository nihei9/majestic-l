package json_log

import (
	"encoding/json"
	"fmt"
	"testing"

	mj "github.com/nihei9/majestic-l"
)

type Log struct {
	Level   string `json:"level" mj:"level"`
	ID      string `json:"id" mj:"id"`
	Message string `json:"message" mj:"message"`
}

func parse(src []byte) (map[string]interface{}, error) {
	log := &Log{}
	json.Unmarshal(src, log)

	return map[string]interface{}{
		"level":   log.Level,
		"id":      log.ID,
		"message": log.Message,
	}, nil
}

func TestPlainTextLog(t *testing.T) {
	config := &mj.Config{
		Parser: parse,
	}
	expectations, err := mj.Expect(
		Log{
			Level:   "info",
			ID:      "I0001",
			Message: "start",
		},
		Log{
			Level:   "info",
			ID:      "I0002",
			Message: "end",
		},
	)
	if err != nil {
		t.Error(err)
	}
	mj.Verify(t, config, expectations, func() {
		b1, _ := json.Marshal(&Log{
			Level:   "info",
			ID:      "I0001",
			Message: "start",
		})
		fmt.Println(string(b1))

		b2, _ := json.Marshal(&Log{
			Level:   "info",
			ID:      "I0002",
			Message: "end",
		})
		fmt.Println(string(b2))
	})
}
