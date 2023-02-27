package audit

import "fmt"

const (
	// List of actions available for audits
	ActionCreate = "created"
	ActionUpdate = "updated"
	ActionDelete = "deleted"

	// List of nouns we recognize for tracking audits
	NounFlag    = "flag"
	NounRule    = "rule"
	NounSegment = "segment"

	// List of emojis recognized
	EmojiCreate = '\U0001F58B'
	EmojiUpdate = '\U0001F58B'
)

// Audit represents a collection of fields that we need for formatting
// the message we will send to the sink
type Audit struct {
	Action       string
	Noun         string
	ResourceName string
	Emoji        int32
}

func (a *Audit) String() string {
	return fmt.Sprintf("%c: The %s `%s` was %s", a.Emoji, a.Noun, a.ResourceName, a.Action)
}

// Auditor sends audit messages to a sink
type Auditor interface {
	SendAudit(*Audit) error
	fmt.Stringer
}
