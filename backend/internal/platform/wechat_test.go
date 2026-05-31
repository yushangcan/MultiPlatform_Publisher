package platform_test

import (
	"context"
	"strings"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestWechatAdapterRewrite(t *testing.T) {
	adapter := platform.NewWechatAdapter()
	draft, err := adapter.Rewrite(context.Background(), content.StructuredContent{
		Topic:          "大学生暑假提升自己",
		Audience:       "大学生",
		CorePoints:     []string{"学习 Go", "完成项目 demo"},
		Keywords:       []string{"大学生", "Go", "项目"},
		SuggestedTitle: "大学生暑假提升自己：一篇适合多平台发布的内容",
	})
	if err != nil {
		t.Fatalf("Rewrite returned error: %v", err)
	}

	if draft.Platform != platform.Wechat {
		t.Fatalf("expected wechat draft, got %s", draft.Platform)
	}
	if !strings.Contains(draft.Body, "## 开篇") {
		t.Fatalf("expected wechat body to contain structured heading, got %q", draft.Body)
	}
	if len(draft.Tags) == 0 {
		t.Fatal("expected tags")
	}
}
