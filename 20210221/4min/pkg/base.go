package pkg

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

type Base struct {
	Msg slack.Msg
}

func (b Base) GetTimestamp() string {
	// msg.Timestamp のフォーマットはこんなやつ `1458422845.000021`
	ts := strings.Split(b.Msg.Timestamp, ".")
	if len(ts) != 2 {
		return "YYYY-MM-DD/HH:MM:SS"
	}
	sec, err := strconv.Atoi(ts[0])
	if err != nil {
		log.Fatalf("strconv.Atoi is failed: %v", err)
		os.Exit(1)
	}
	msec, err := strconv.Atoi(ts[1])
	if err != nil {
		log.Fatalf("strconv.Atoi is failed: %v", err)
		os.Exit(1)
	}
	t := time.Unix(int64(sec), int64(msec))
	return t.Format(`2006-01-02/03:04:05`)
}
