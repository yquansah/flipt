package slack

import (
	"context"
	"time"

	"github.com/slack-go/slack"
	"go.flipt.io/flipt/internal/server/audit"
)

const auditorType = "slack"

var _ audit.Auditor = (*SlackAuditor)(nil)

// SlackAuditor is an implementation of an Auditor that will send messages
// to slack as a sink
type SlackAuditor struct {
	slackClient *slack.Client
	channelName string
}

// NewSlackAuditor is the constructor for a SlackAuditor
func NewSlackAuditor(apiToken string, channelName string) *SlackAuditor {
	slackClient := slack.New(apiToken)

	return &SlackAuditor{
		slackClient: slackClient,
		channelName: channelName,
	}
}

// SendAudit sends an Audit message to slack
func (s *SlackAuditor) SendAudit(audit *audit.Audit) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Send message over to slack. We only care about the error here as the first
	// two return values are not useful for our purposes
	_, _, err := s.slackClient.PostMessageContext(ctx, s.channelName, slack.MsgOptionText(audit.String(), false))

	return err
}

func (s *SlackAuditor) String() string {
	return auditorType
}
