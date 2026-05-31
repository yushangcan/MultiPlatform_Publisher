package llm_test

import (
	"errors"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/config"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/llm"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestNewProviderDefaultsToRule(t *testing.T) {
	registry, err := platform.NewRegistry(platform.NewWechatAdapter())
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	provider, err := llm.NewProvider(config.Config{}, analyzer.NewRuleAnalyzer(), registry)
	if err != nil {
		t.Fatalf("NewProvider returned error: %v", err)
	}

	if provider.Name() != "rule" {
		t.Fatalf("expected rule provider, got %s", provider.Name())
	}
}

func TestNewProviderRequiresDeepSeekAPIKey(t *testing.T) {
	registry, err := platform.NewRegistry(platform.NewWechatAdapter())
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	_, err = llm.NewProvider(config.Config{LLMProvider: "deepseek"}, analyzer.NewRuleAnalyzer(), registry)
	if !errors.Is(err, llm.ErrMissingAPIKey) {
		t.Fatalf("expected ErrMissingAPIKey, got %v", err)
	}
}
