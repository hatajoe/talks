package dialect

import (
	"strings"

	"github.com/hatajoe/talks/20210221/4min/pkg"
	"github.com/slack-go/slack"
)

const (
	coverallsBotID = `B0U0D9473`
)

type CoverallsMatcher struct{}

func (m CoverallsMatcher) IsMatch(msg slack.Msg) bool {
	return msg.BotID == coverallsBotID
}

func (m CoverallsMatcher) NewNotification(msg slack.Msg) pkg.Notification {
	return &CoverallsNotification{
		Base: pkg.Base{
			Msg: msg,
		},
	}
}

type CoverallsNotification struct {
	pkg.Base
}

func (n *CoverallsNotification) GetServiceName() string {
	return "Coveralls"
}

func (n *CoverallsNotification) GetMessage() string {
	if len(n.Base.Msg.Attachments) < 1 {
		return "unknow message"
	}
	return strings.Replace(n.Base.Msg.Attachments[0].Text, "\\n", "\n", -1)
}
