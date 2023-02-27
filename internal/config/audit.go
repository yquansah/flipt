package config

import (
	"encoding/json"

	"github.com/spf13/viper"
)

// AuditConfig contains fields, which enable and configure
// Flipt's various audit lgogging mechanisms.
//
// Currently, flipt support slack for audit logs.
type AuditConfig struct {
	Enabled bool          `json:"enabled" mapstructure:"enabled"`
	Auditor AuditorSink   `json:"auditor,omitempty" mapstructure:"auditor"`
	Slack   SlackConfig   `json:"slack,omitempty" mapstructure:"slack"`
	Discord DiscordConfig `json:"discord,omitempty" mapstructure:"discord"`
}

func (c *AuditConfig) setDefaults(v *viper.Viper) {
	v.SetDefault("audit", map[string]any{
		"enabled": false,
		"auditor": AuditorSlack,
		"slack": map[string]any{
			"apiToken": "myapitoken",
			"channel":  "foobar",
		},
		"discord": map[string]any{
			"botToken":  "mybottoken",
			"channelID": "channelID",
		},
	})
}

// AuditorSink is either discord or slack as a sink
type AuditorSink uint8

func (c AuditorSink) String() string {
	return auditorSinkToString[c]
}

func (c AuditorSink) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

const (
	_ AuditorSink = iota
	// AuditorSlack ...
	AuditorSlack
	// AuditorDiscord
	AuditorDiscord
)

var (
	auditorSinkToString = map[AuditorSink]string{
		AuditorSlack:   "slack",
		AuditorDiscord: "discord",
	}

	stringToAuditorSink = map[string]AuditorSink{
		"slack":   AuditorSlack,
		"discord": AuditorDiscord,
	}
)

// SlackConfig contains configuration for communicating with slack
type SlackConfig struct {
	ApiToken string `json:"apiToken,omitempty" mapstructure:"apiToken"`
	Channel  string `json:"channel,omitempty" mapstructure:"channel"`
}

// DiscordConfig contains configuration for communicating with discord
type DiscordConfig struct {
	BotToken  string `json:"botToken,omitempty" mapstructure:"botToken"`
	ChannelID string `json:"channelID,omitempty" mapstructure:"channelID"`
}
