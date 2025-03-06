package slack

// Block represents a Slack block element
type Block struct {
	Type string       `json:"type"`
	Text *TextObject  `json:"text,omitempty"`
}

// TextObject represents text within a block
type TextObject struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}
