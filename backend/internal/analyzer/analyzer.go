package analyzer

import (
	"context"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

type ContentAnalyzer interface {
	Analyze(ctx context.Context, input content.RawInput) (content.StructuredContent, error)
}
