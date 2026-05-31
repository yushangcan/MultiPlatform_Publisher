package rewrite_test

import (
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/rewrite"
)

func TestRewriteRequestValidateAcceptsCompleteRequest(t *testing.T) {
	request := rewrite.RewriteRequest{
		Content: content.StructuredContent{
			Topic:      "大学生暑假提升",
			CorePoints: []string{"学习 Go", "完成项目"},
		},
		Platforms: []platform.Platform{platform.Wechat, platform.Zhihu},
	}

	if issues := request.Validate(); len(issues) != 0 {
		t.Fatalf("expected no validation issues, got %d", len(issues))
	}
}

func TestRewriteRequestValidateRejectsMissingFields(t *testing.T) {
	request := rewrite.RewriteRequest{}

	issues := request.Validate()

	if len(issues) != 2 {
		t.Fatalf("expected 2 validation issues, got %d", len(issues))
	}
}

func TestRewriteResponseDraftFor(t *testing.T) {
	response := rewrite.RewriteResponse{
		Drafts: []platform.PlatformDraft{
			{Platform: platform.Wechat, Title: "公众号标题"},
			{Platform: platform.Zhihu, Title: "知乎标题"},
		},
	}

	draft, ok := response.DraftFor(platform.Zhihu)
	if !ok {
		t.Fatal("expected zhihu draft to exist")
	}
	if draft.Title != "知乎标题" {
		t.Fatalf("expected zhihu title, got %q", draft.Title)
	}
}
