package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/hatajoe/talks/20210221/4min/pkg"
)

type SlackAPIResponse struct {
	OK       bool        `json:"ok"`
	Messages []slack.Msg `json:"messages"`
}

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

	p := pkg.NewNotificationParser([]pkg.Matcher{})
	notifications := []pkg.Notification{}
	for _, msg := range res.Messages {
		n, err := p.Parse(msg)
		if err != nil {
			continue
		}
		notifications = append(notifications, n)
	}

	for _, n := range notifications {
		fmt.Printf("%v %v %v\n", n.GetTimestamp(), n.GetServiceName(), n.GetMessage())
	}
}
