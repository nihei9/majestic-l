package text_log

import (
	"fmt"
	"regexp"
	"testing"

	mj "github.com/nihei9/majestic-l"
)

type Log struct {
	Level   string `mj:"level"`
	ID      string `mj:"id"`
	Message string `mj:"message"`
}

func parse(src []byte) (map[string]interface{}, error) {
	re := regexp.MustCompile(`^\s*(.+)\s+(.+)\s+(.+)$`)
	s := re.FindSubmatch(src)
	if len(s) != 4 {
		return nil, fmt.Errorf("Pase error")
	}

	return map[string]interface{}{
		"level":   string(s[1]),
		"id":      string(s[2]),
		"message": string(s[3]),
	}, nil
}

func TestTextLog(t *testing.T) {
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
		fmt.Println("info I0001 start")
		fmt.Println("info I0002 end")
	})
}
