package llm

import (
	"context"
	"fmt"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

type RuleProvider struct {
	analyzer analyzer.ContentAnalyzer
	registry *platform.Registry
}

func NewRuleProvider(analyzer analyzer.ContentAnalyzer, registry *platform.Registry) RuleProvider {
	return RuleProvider{
		analyzer: analyzer,
		registry: registry,
	}
}

func (provider RuleProvider) Name() string {
	return "rule"
}

func (provider RuleProvider) Analyze(ctx context.Context, input content.RawInput) (content.StructuredContent, error) {
	return provider.analyzer.Analyze(ctx, input)
}

func (provider RuleProvider) Rewrite(ctx context.Context, structured content.StructuredContent, target platform.Platform) (platform.PlatformDraft, error) {
	adapter, err := provider.registry.MustGet(target)
	if err != nil {
		return platform.PlatformDraft{}, fmt.Errorf("get platform adapter: %w", err)
	}
	return adapter.Rewrite(ctx, structured)
}
