package llm_test

import (
	"context"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/llm"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestRuleProviderAnalyzeAndRewrite(t *testing.T) {
	registry, err := platform.NewRegistry(platform.NewWechatAdapter())
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	provider := llm.NewRuleProvider(analyzer.NewRuleAnalyzer(), registry)
	structured, err := provider.Analyze(context.Background(), content.RawInput{
		Text: "我想写一篇关于大学生暑假提升自己的内容，主要讲学习 Go 和做项目。",
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	draft, err := provider.Rewrite(context.Background(), structured, platform.Wechat)
	if err != nil {
		t.Fatalf("Rewrite returned error: %v", err)
	}

	if draft.Platform != platform.Wechat {
		t.Fatalf("expected wechat draft, got %s", draft.Platform)
	}
}
