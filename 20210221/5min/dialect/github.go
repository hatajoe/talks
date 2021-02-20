package dialect

import (
	"strings"

	"github.com/hatajoe/talks/20210221/4min/pkg"
	"github.com/slack-go/slack"
)

const (
	githubBotID = `B0U0BK8TA`
)

type GitHubMatcher struct{}

func (m GitHubMatcher) IsMatch(msg slack.Msg) bool {
	return msg.BotID == githubBotID
}

func (m GitHubMatcher) NewNotification(msg slack.Msg) pkg.Notification {
	return &GitHubNotification{
		Base: pkg.Base{
			Msg: msg,
		},
	}
}

type GitHubNotification struct {
	pkg.Base
}

func (n *GitHubNotification) GetServiceName() string {
	return "GitHub"
}

func (n *GitHubNotification) GetMessage() string {
	if len(n.Base.Msg.Attachments) < 1 {
		return "unknow message"
	}
	return strings.Replace(n.Base.Msg.Attachments[0].Text, "\\n", "\n", -1)
}
