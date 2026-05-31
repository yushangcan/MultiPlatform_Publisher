package platform_test

import (
	"context"
	"strings"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestZhihuAdapterRewrite(t *testing.T) {
	adapter := platform.NewZhihuAdapter()
	draft, err := adapter.Rewrite(context.Background(), content.StructuredContent{
		Topic:          "大学生暑假提升自己",
		Audience:       "大学生",
		CorePoints:     []string{"学习 Go", "完成项目 demo"},
		Keywords:       []string{"大学生", "Go", "项目"},
		SuggestedTitle: "大学生暑假提升自己应该怎么做？",
	})
	if err != nil {
		t.Fatalf("Rewrite returned error: %v", err)
	}

	if draft.Platform != platform.Zhihu {
		t.Fatalf("expected zhihu draft, got %s", draft.Platform)
	}
	if !strings.Contains(draft.Body, "这个问题") {
		t.Fatalf("expected zhihu body to use answer style, got %q", draft.Body)
	}
	if len(draft.Tags) == 0 {
		t.Fatal("expected tags")
	}
}
