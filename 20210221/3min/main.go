package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

var (
	botID = map[string]string{
		"B0TVCS8FN": "circleci",
		"B0U0BK8TA": "github",
		"B0U0D9473": "coveralls",
	}
)

type SlackAPIResponse struct {
	OK       bool        `json:"ok"`
	Messages []slack.Msg `json:"messages"`
}

type Notification struct {
	Timestamp   string
	ServiceName string
	Repository  string
	Message     string
}

func (n Notification) String() string {
	return fmt.Sprintf("%v %v %v %v", n.Timestamp, n.ServiceName, n.Repository, n.Message)
}

type Notifications []*Notification

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("ioutil.ReadAll is failed: %v", err)
		os.Exit(1)
	}

	res := SlackAPIResponse{}
	if err := json.Unmarshal([]byte(buf), &res); err != nil {
		log.Fatalf("json.Unmarshal is failed: %v", err)
		os.Exit(1)
	}

	notifications := []*Notification{}
	for _, msg := range res.Messages {
		// botID マッチしないやつは無視
		service, ok := botID[msg.BotID]
		if !ok {
			continue
		}
		// msg.Timestamp のフォーマットはこんなやつ `1458422845.000021`
		ts := strings.Split(msg.Timestamp, ".")
		if len(ts) != 2 {
			continue
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

		notifications = append(notifications, &Notification{
			Timestamp:   t.Format(`2006-01-02/03:04:05`),
			ServiceName: service,
			Repository:  "",
			Message:     "",
		})
	}

	for _, n := range notifications {
		fmt.Println(n)
	}
}
