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
	Enabled bool        `json:"enabled" mapstructure:"enabled"`
	Auditor AuditorSink `json:"auditor,omitempty" mapstructure:"auditor"`
	Slack   SlackConfig `json:"slack,omitempty" mapstructure:"slack"`
}

func (c *AuditConfig) setDefaults(v *viper.Viper) {
	v.SetDefault("audit", map[string]any{
		"enabled": false,
		"auditor": AuditorSlack,
		"slack": map[string]any{
			"apiToken": "myapitoken",
			"channel":  "foobar",
		},
	})
}

// AuditorSink is slack for now
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
)

var (
	auditorSinkToString = map[AuditorSink]string{
		AuditorSlack: "slack",
	}

	stringToAuditorSink = map[string]AuditorSink{
		"slack": AuditorSlack,
	}
)

// SlackConfig contains information which is useful for communicating through to slack
type SlackConfig struct {
	ApiToken string `json:"apiToken,omitempty" mapstructure:"apiToken"`
	Channel  string `json:"channel,omitempty" mapstructure:"channel"`
}
