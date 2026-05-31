package platform

import (
	"context"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

type Adapter interface {
	Platform() Platform
	Rewrite(ctx context.Context, content content.StructuredContent) (PlatformDraft, error)
	Validate(draft PlatformDraft) []ValidationIssue
}

type AdapterInfo struct {
	Platform    Platform `json:"platform"`
	DisplayName string   `json:"display_name"`
}

func NewAdapterInfo(platform Platform) AdapterInfo {
	return AdapterInfo{
		Platform:    platform,
		DisplayName: platform.DisplayName(),
	}
}
