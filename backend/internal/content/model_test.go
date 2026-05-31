package content_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

func TestRawInputValidateRequiresText(t *testing.T) {
	input := content.RawInput{Text: "   "}

	err := input.Validate()

	if !errors.Is(err, content.ErrRawInputTextRequired) {
		t.Fatalf("expected ErrRawInputTextRequired, got %v", err)
	}
}

func TestRawInputValidateRejectsLongText(t *testing.T) {
	input := content.RawInput{Text: strings.Repeat("a", content.MaxRawInputTextLength+1)}

	err := input.Validate()

	if !errors.Is(err, content.ErrRawInputTextTooLong) {
		t.Fatalf("expected ErrRawInputTextTooLong, got %v", err)
	}
}

func TestStructuredContentHasCore(t *testing.T) {
	structured := content.StructuredContent{
		Topic:      "大学生暑假提升",
		CorePoints: []string{"学习 Go", "  ", "完成项目 demo"},
	}

	if !structured.HasCore() {
		t.Fatal("expected structured content to have core information")
	}

	points := structured.NonEmptyCorePoints()
	if len(points) != 2 {
		t.Fatalf("expected 2 non-empty core points, got %d", len(points))
	}
}
