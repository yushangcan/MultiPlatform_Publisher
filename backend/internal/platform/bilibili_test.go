package platform_test

import (
	"context"
	"strings"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestBilibiliAdapterRewrite(t *testing.T) {
	adapter := platform.NewBilibiliAdapter()
	draft, err := adapter.Rewrite(context.Background(), content.StructuredContent{
		Topic:      "大学生暑假提升自己",
		CorePoints: []string{"学习 Go", "完成项目 demo"},
		Keywords:   []string{"大学生", "Go", "项目"},
	})
	if err != nil {
		t.Fatalf("Rewrite returned error: %v", err)
	}

	if draft.Platform != platform.Bilibili {
		t.Fatalf("expected bilibili draft, got %s", draft.Platform)
	}
	if !strings.Contains(draft.Body, "评论区") {
		t.Fatalf("expected bilibili body to include interaction guidance, got %q", draft.Body)
	}
	if len(adapter.Validate(draft)) != 0 {
		t.Fatal("expected valid bilibili draft")
	}
}
