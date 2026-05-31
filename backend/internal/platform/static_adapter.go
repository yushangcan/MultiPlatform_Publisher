package platform

import (
	"context"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

type StaticAdapter struct {
	platform Platform
}

func NewStaticAdapter(platform Platform) StaticAdapter {
	return StaticAdapter{platform: platform}
}

func (adapter StaticAdapter) Platform() Platform {
	return adapter.platform
}

func (adapter StaticAdapter) Rewrite(_ context.Context, structured content.StructuredContent) (PlatformDraft, error) {
	title := structured.SuggestedTitle
	if title == "" {
		title = structured.Topic
	}

	return PlatformDraft{
		Platform: adapter.platform,
		Title:    title,
		Body:     structured.Topic,
		Tags:     structured.NonEmptyKeywords(),
	}, nil
}

func (adapter StaticAdapter) Validate(draft PlatformDraft) []ValidationIssue {
	return draft.ValidateBasic()
}
