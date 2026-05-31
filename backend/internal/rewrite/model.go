package rewrite

import (
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

type RewriteRequest struct {
	Content   content.StructuredContent `json:"content"`
	Platforms []platform.Platform       `json:"platforms"`
}

func (request RewriteRequest) Validate() []platform.ValidationIssue {
	issues := make([]platform.ValidationIssue, 0)
	if !request.Content.HasCore() {
		issues = append(issues, platform.ValidationIssue{
			Field:    "content",
			Code:     "content_core_required",
			Message:  "structured content requires a topic and at least one core point",
			Severity: platform.SeverityError,
		})
	}
	if len(request.Platforms) == 0 {
		issues = append(issues, platform.ValidationIssue{
			Field:    "platforms",
			Code:     "platforms_required",
			Message:  "at least one target platform is required",
			Severity: platform.SeverityError,
		})
	}
	for _, item := range request.Platforms {
		if !item.IsSupported() {
			issues = append(issues, platform.ValidationIssue{
				Field:    "platforms",
				Code:     "unsupported_platform",
				Message:  "platform is not supported",
				Severity: platform.SeverityError,
			})
		}
	}
	return issues
}

type RewriteResponse struct {
	Content content.StructuredContent  `json:"content"`
	Drafts  []platform.PlatformDraft   `json:"drafts"`
	Issues  []platform.ValidationIssue `json:"issues,omitempty"`
}

func (response RewriteResponse) DraftFor(target platform.Platform) (platform.PlatformDraft, bool) {
	for _, draft := range response.Drafts {
		if draft.Platform == target {
			return draft, true
		}
	}
	return platform.PlatformDraft{}, false
}
