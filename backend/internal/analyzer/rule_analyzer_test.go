package analyzer_test

import (
	"context"
	"errors"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

func TestRuleAnalyzerAnalyzeExtractsStructuredContent(t *testing.T) {
	service := analyzer.NewRuleAnalyzer()
	input := content.RawInput{
		Text: "我想写一篇关于大学生暑假提升自己的内容，主要讲学习 Go、做项目、准备简历和保持运动，语气希望轻松一点。",
	}

	result, err := service.Analyze(context.Background(), input)
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if result.Topic == "" {
		t.Fatal("expected topic to be inferred")
	}
	if result.Audience != "大学生" {
		t.Fatalf("expected audience 大学生, got %q", result.Audience)
	}
	if result.Tone != "轻松、实用" {
		t.Fatalf("expected light tone, got %q", result.Tone)
	}
	if len(result.CorePoints) == 0 {
		t.Fatal("expected core points")
	}
	if len(result.Keywords) == 0 {
		t.Fatal("expected keywords")
	}
	if result.SuggestedTitle == "" {
		t.Fatal("expected suggested title")
	}
}

func TestRuleAnalyzerAnalyzeUsesRequestedContentType(t *testing.T) {
	service := analyzer.NewRuleAnalyzer()
	input := content.RawInput{
		Text:        "整理一个 Go 项目实训复盘，讲项目背景、实现过程和收获。",
		ContentType: "项目复盘",
	}

	result, err := service.Analyze(context.Background(), input)
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if result.ContentType != "项目复盘" {
		t.Fatalf("expected requested content type, got %q", result.ContentType)
	}
}

func TestRuleAnalyzerAnalyzeRejectsInvalidInput(t *testing.T) {
	service := analyzer.NewRuleAnalyzer()
	_, err := service.Analyze(context.Background(), content.RawInput{Text: " "})

	if !errors.Is(err, content.ErrRawInputTextRequired) {
		t.Fatalf("expected ErrRawInputTextRequired, got %v", err)
	}
}
