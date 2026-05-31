package platform_test

import (
	"context"
	"strings"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestXiaohongshuAdapterRewrite(t *testing.T) {
	adapter := platform.NewXiaohongshuAdapter()
	draft, err := adapter.Rewrite(context.Background(), content.StructuredContent{
		Topic:      "大学生暑假提升自己",
		CorePoints: []string{"学习 Go", "完成项目 demo"},
		Keywords:   []string{"大学生", "Go", "项目"},
	})
	if err != nil {
		t.Fatalf("Rewrite returned error: %v", err)
	}

	if draft.Platform != platform.Xiaohongshu {
		t.Fatalf("expected xiaohongshu draft, got %s", draft.Platform)
	}
	if !strings.Contains(draft.Body, "#大学生") {
		t.Fatalf("expected xiaohongshu body to include hashtags, got %q", draft.Body)
	}
	if len(adapter.Validate(draft)) != 0 {
		t.Fatal("expected valid xiaohongshu draft")
	}
}
