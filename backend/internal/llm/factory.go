package llm

import (
	"fmt"
	"strings"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/config"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func NewProvider(cfg config.Config, analyzer analyzer.ContentAnalyzer, registry *platform.Registry) (Provider, error) {
	switch strings.ToLower(strings.TrimSpace(cfg.LLMProvider)) {
	case "", "rule":
		return NewRuleProvider(analyzer, registry), nil
	case "deepseek":
		return NewDeepSeekProvider(cfg)
	default:
		return nil, fmt.Errorf("unsupported llm provider: %s", cfg.LLMProvider)
	}
}
