package rewrite_test

import (
	"context"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/rewrite"
)

func TestServiceRewrite(t *testing.T) {
	registry, err := platform.NewRegistry(
		platform.NewWechatAdapter(),
		platform.NewZhihuAdapter(),
		platform.NewBilibiliAdapter(),
		platform.NewXiaohongshuAdapter(),
	)
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	service := rewrite.NewService(registry)
	response, issues := service.Rewrite(context.Background(), rewrite.RewriteRequest{
		Content: content.StructuredContent{
			Topic:      "大学生暑假提升自己",
			CorePoints: []string{"学习 Go", "完成项目 demo"},
		},
		Platforms: []platform.Platform{platform.Wechat, platform.Zhihu, platform.Bilibili, platform.Xiaohongshu},
	})

	if len(issues) != 0 {
		t.Fatalf("expected no issues, got %d", len(issues))
	}
	if len(response.Drafts) != 4 {
		t.Fatalf("expected 4 drafts, got %d", len(response.Drafts))
	}
	if _, ok := response.DraftFor(platform.Wechat); !ok {
		t.Fatal("expected wechat draft")
	}
	if _, ok := response.DraftFor(platform.Xiaohongshu); !ok {
		t.Fatal("expected xiaohongshu draft")
	}
}

func TestServiceRewriteRejectsInvalidRequest(t *testing.T) {
	registry, err := platform.NewRegistry(platform.NewWechatAdapter())
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	service := rewrite.NewService(registry)
	_, issues := service.Rewrite(context.Background(), rewrite.RewriteRequest{})

	if len(issues) == 0 {
		t.Fatal("expected validation issues")
	}
}
