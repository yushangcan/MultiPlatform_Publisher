package platform_test

import (
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestSupportedPlatformsAreStable(t *testing.T) {
	platforms := platform.SupportedPlatforms()

	expected := []platform.Platform{
		platform.Wechat,
		platform.Zhihu,
		platform.Bilibili,
		platform.Xiaohongshu,
	}

	if len(platforms) != len(expected) {
		t.Fatalf("expected %d platforms, got %d", len(expected), len(platforms))
	}
	for i, item := range expected {
		if platforms[i] != item {
			t.Fatalf("expected platform %s at %d, got %s", item, i, platforms[i])
		}
	}
}

func TestPlatformDraftValidateBasic(t *testing.T) {
	draft := platform.PlatformDraft{
		Platform: platform.Platform("unknown"),
		Title:    " ",
		Body:     "",
	}

	issues := draft.ValidateBasic()

	if len(issues) != 3 {
		t.Fatalf("expected 3 issues, got %d", len(issues))
	}
	for _, issue := range issues {
		if issue.Severity != platform.SeverityError {
			t.Fatalf("expected error severity, got %s", issue.Severity)
		}
	}
}
