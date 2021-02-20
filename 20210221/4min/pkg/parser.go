package pkg

import (
	"fmt"

	"github.com/slack-go/slack"
)

type Notification interface {
	GetTimestamp() string
	GetServiceName() string
	GetMessage() string
}

type Matcher interface {
	IsMatch(msg slack.Msg) bool
	NewNotification(msg slack.Msg) Notification
}

type NotificationParser struct {
	matcher []Matcher
}

func NewNotificationParser(matchers []Matcher) *NotificationParser {
	return &NotificationParser{
		matcher: matchers,
	}
}

func (f NotificationParser) Parse(msg slack.Msg) (Notification, error) {
	for _, m := range f.matcher {
		if m.IsMatch(msg) {
			return m.NewNotification(msg), nil
		}
	}
	return nil, fmt.Errorf("msg is not supported by current matchers: %v", msg.BotID)
}
