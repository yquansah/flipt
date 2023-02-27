package discord

import (
	"github.com/bwmarrin/discordgo"
	"go.flipt.io/flipt/internal/server/audit"
)

const auditorType = "discord"

// DiscordAuditor is the discord implementation of an auditor that is configured to send Audit messages
// to discord as a sink
type DiscordAuditor struct {
	session   *discordgo.Session
	channelID string
}

// NewDiscordAuditor is the constructor for a DiscordAuditor
func NewDiscordAuditor(botToken, channelID string) *DiscordAuditor {
	session, _ := discordgo.New("Bot " + botToken)

	return &DiscordAuditor{
		session:   session,
		channelID: channelID,
	}
}

// SendAudit sends an Audit message to Discord
func (d *DiscordAuditor) SendAudit(audit *audit.Audit) error {
	_, err := d.session.ChannelMessageSend(d.channelID, audit.String())
	return err
}

func (d *DiscordAuditor) String() string {
	return auditorType
}
