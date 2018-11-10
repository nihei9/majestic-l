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

func TestPlainTextLog(t *testing.T) {
	config := &mj.Config{
		Parser: mj.JSONParser,
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
