package llm

import (
	"context"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

type Provider interface {
	Name() string
	Analyze(ctx context.Context, input content.RawInput) (content.StructuredContent, error)
	Rewrite(ctx context.Context, structured content.StructuredContent, target platform.Platform) (platform.PlatformDraft, error)
}
