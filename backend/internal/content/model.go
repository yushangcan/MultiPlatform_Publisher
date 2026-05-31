package content

import (
	"errors"
	"strings"
	"time"
)

const MaxRawInputTextLength = 20000

var (
	ErrRawInputTextRequired = errors.New("raw input text is required")
	ErrRawInputTextTooLong  = errors.New("raw input text exceeds maximum length")
)

type RawInput struct {
	ID          string    `json:"id,omitempty"`
	Text        string    `json:"text"`
	ContentType string    `json:"content_type,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (input RawInput) NormalizedText() string {
	return strings.TrimSpace(input.Text)
}

func (input RawInput) Validate() error {
	text := input.NormalizedText()
	if text == "" {
		return ErrRawInputTextRequired
	}
	if len([]rune(text)) > MaxRawInputTextLength {
		return ErrRawInputTextTooLong
	}
	return nil
}

type StructuredContent struct {
	ID             string    `json:"id,omitempty"`
	SourceID       string    `json:"source_id,omitempty"`
	Topic          string    `json:"topic"`
	Audience       string    `json:"audience,omitempty"`
	ContentType    string    `json:"content_type,omitempty"`
	Tone           string    `json:"tone,omitempty"`
	CorePoints     []string  `json:"core_points"`
	Keywords       []string  `json:"keywords,omitempty"`
	SuggestedTitle string    `json:"suggested_title,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

func (content StructuredContent) HasCore() bool {
	return strings.TrimSpace(content.Topic) != "" && len(content.NonEmptyCorePoints()) > 0
}

func (content StructuredContent) NonEmptyCorePoints() []string {
	return nonEmptyStrings(content.CorePoints)
}

func (content StructuredContent) NonEmptyKeywords() []string {
	return nonEmptyStrings(content.Keywords)
}

func nonEmptyStrings(values []string) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}
