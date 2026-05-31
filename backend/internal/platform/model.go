package platform

import (
	"strings"
	"time"
)

type Platform string

const (
	Wechat      Platform = "wechat"
	Zhihu       Platform = "zhihu"
	Bilibili    Platform = "bilibili"
	Xiaohongshu Platform = "xiaohongshu"
)

var displayNames = map[Platform]string{
	Wechat:      "微信公众号",
	Zhihu:       "知乎",
	Bilibili:    "B站",
	Xiaohongshu: "小红书",
}

func SupportedPlatforms() []Platform {
	return []Platform{Wechat, Zhihu, Bilibili, Xiaohongshu}
}

func (platform Platform) IsSupported() bool {
	_, ok := displayNames[platform]
	return ok
}

func (platform Platform) DisplayName() string {
	if name, ok := displayNames[platform]; ok {
		return name
	}
	return string(platform)
}

type ValidationSeverity string

const (
	SeverityWarning ValidationSeverity = "warning"
	SeverityError   ValidationSeverity = "error"
)

type ValidationIssue struct {
	Field    string             `json:"field"`
	Code     string             `json:"code"`
	Message  string             `json:"message"`
	Severity ValidationSeverity `json:"severity"`
}

type PlatformDraft struct {
	ID        string            `json:"id,omitempty"`
	Platform  Platform          `json:"platform"`
	Title     string            `json:"title"`
	Body      string            `json:"body"`
	Tags      []string          `json:"tags,omitempty"`
	Warnings  []ValidationIssue `json:"warnings,omitempty"`
	CreatedAt time.Time         `json:"created_at,omitempty"`
}

func (draft PlatformDraft) ValidateBasic() []ValidationIssue {
	issues := make([]ValidationIssue, 0, 3)
	if !draft.Platform.IsSupported() {
		issues = append(issues, ValidationIssue{
			Field:    "platform",
			Code:     "unsupported_platform",
			Message:  "platform is not supported",
			Severity: SeverityError,
		})
	}
	if strings.TrimSpace(draft.Title) == "" {
		issues = append(issues, ValidationIssue{
			Field:    "title",
			Code:     "title_required",
			Message:  "title is required",
			Severity: SeverityError,
		})
	}
	if strings.TrimSpace(draft.Body) == "" {
		issues = append(issues, ValidationIssue{
			Field:    "body",
			Code:     "body_required",
			Message:  "body is required",
			Severity: SeverityError,
		})
	}
	return issues
}
