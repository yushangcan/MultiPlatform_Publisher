package config_test

import (
	"testing"
	"time"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/config"
)

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("LLM_PROVIDER", "")
	t.Setenv("LLM_API_KEY", "")
	t.Setenv("LLM_MODEL", "")
	t.Setenv("LLM_BASE_URL", "")
	t.Setenv("LLM_TIMEOUT_SECONDS", "")

	cfg := config.Load()

	if cfg.Port != config.DefaultPort {
		t.Fatalf("expected default port, got %q", cfg.Port)
	}
	if cfg.LLMProvider != config.DefaultLLMProvider {
		t.Fatalf("expected default provider, got %q", cfg.LLMProvider)
	}
	if cfg.LLMTimeout != config.DefaultLLMTimeout {
		t.Fatalf("expected default timeout, got %s", cfg.LLMTimeout)
	}
}

func TestLoadReadsEnvironment(t *testing.T) {
	t.Setenv("PORT", "18080")
	t.Setenv("LLM_PROVIDER", "deepseek")
	t.Setenv("LLM_API_KEY", "test-key")
	t.Setenv("LLM_MODEL", "test-model")
	t.Setenv("LLM_BASE_URL", "https://example.com")
	t.Setenv("LLM_TIMEOUT_SECONDS", "5")

	cfg := config.Load()

	if cfg.Port != "18080" {
		t.Fatalf("expected custom port, got %q", cfg.Port)
	}
	if cfg.LLMProvider != "deepseek" {
		t.Fatalf("expected deepseek provider, got %q", cfg.LLMProvider)
	}
	if cfg.LLMAPIKey != "test-key" {
		t.Fatal("expected API key from environment")
	}
	if cfg.LLMTimeout != 5*time.Second {
		t.Fatalf("expected 5s timeout, got %s", cfg.LLMTimeout)
	}
}
