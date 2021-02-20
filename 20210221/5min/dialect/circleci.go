package dialect

import (
	"strings"

	"github.com/hatajoe/talks/20210221/4min/pkg"
	"github.com/slack-go/slack"
)

const (
	circleciBotID = `B0TVCS8FN`
)

type CircleCIMatcher struct{}

func (m CircleCIMatcher) IsMatch(msg slack.Msg) bool {
	return msg.BotID == circleciBotID
}

func (m CircleCIMatcher) NewNotification(msg slack.Msg) pkg.Notification {
	return &CircleCINotification{
		Base: pkg.Base{
			Msg: msg,
		},
	}
}

type CircleCINotification struct {
	pkg.Base
}

func (n *CircleCINotification) GetServiceName() string {
	return "CircleCI"
}

func (n *CircleCINotification) GetMessage() string {
	if len(n.Base.Msg.Attachments) < 1 {
		return "unknow message"
	}
	return strings.Replace(n.Base.Msg.Attachments[0].Text, "\n", "", -1)
}
