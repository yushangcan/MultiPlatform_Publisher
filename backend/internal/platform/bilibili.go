package platform

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

type BilibiliAdapter struct{}

func NewBilibiliAdapter() BilibiliAdapter {
	return BilibiliAdapter{}
}

func (adapter BilibiliAdapter) Platform() Platform {
	return Bilibili
}

func (adapter BilibiliAdapter) Rewrite(_ context.Context, structured content.StructuredContent) (PlatformDraft, error) {
	topic := topicOrFallback(structured)
	title := titleOrTopic(structured)
	if title == "" {
		title = topic + "｜适合做成一期分享"
	}

	body := fmt.Sprintf(
		"这期想和大家聊聊：%s。\n\n我会重点讲这几个部分：\n%s\n\n如果你也在做类似准备，可以先收藏，后面按这几个方向一步步推进。评论区也欢迎补充你的做法。",
		topic,
		bulletPoints(structured.NonEmptyCorePoints(), "- "),
	)

	return PlatformDraft{
		Platform:  Bilibili,
		Title:     title,
		Body:      body,
		Tags:      joinTags(structured.NonEmptyKeywords(), topic),
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (adapter BilibiliAdapter) Validate(draft PlatformDraft) []ValidationIssue {
	issues := draft.ValidateBasic()
	if len([]rune(draft.Title)) > 80 {
		issues = append(issues, ValidationIssue{
			Field:    "title",
			Code:     "bilibili_title_too_long",
			Message:  "B站标题建议控制在 80 个字以内，便于列表页展示",
			Severity: SeverityWarning,
		})
	}
	if !strings.Contains(draft.Body, "评论") {
		issues = append(issues, ValidationIssue{
			Field:    "body",
			Code:     "bilibili_interaction_missing",
			Message:  "B站内容建议加入评论区互动引导",
			Severity: SeverityWarning,
		})
	}
	return issues
}
