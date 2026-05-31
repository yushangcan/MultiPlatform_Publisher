package platform

import (
	"context"
	"fmt"
	"time"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

type WechatAdapter struct{}

func NewWechatAdapter() WechatAdapter {
	return WechatAdapter{}
}

func (adapter WechatAdapter) Platform() Platform {
	return Wechat
}

func (adapter WechatAdapter) Rewrite(_ context.Context, structured content.StructuredContent) (PlatformDraft, error) {
	topic := topicOrFallback(structured)
	title := titleOrTopic(structured)
	if title == "" {
		title = topic + "：一篇结构化分享"
	}

	body := fmt.Sprintf(
		"## 开篇\n%s 是一个值得系统梳理的话题。面向 %s，这篇内容会尽量用清晰的结构说明核心思路。\n\n## 核心内容\n%s\n\n## 总结\n如果只记住一点，就是先明确目标，再把行动拆成可以执行的小步骤。这样内容不仅更完整，也更适合在公众号中沉淀为长期可读的文章。",
		topic,
		audienceOrFallback(structured),
		numberedPoints(structured.NonEmptyCorePoints()),
	)

	return PlatformDraft{
		Platform:  Wechat,
		Title:     title,
		Body:      body,
		Tags:      joinTags(structured.NonEmptyKeywords(), topic),
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (adapter WechatAdapter) Validate(draft PlatformDraft) []ValidationIssue {
	issues := draft.ValidateBasic()
	if len([]rune(draft.Title)) > 64 {
		issues = append(issues, ValidationIssue{
			Field:    "title",
			Code:     "wechat_title_too_long",
			Message:  "公众号标题建议控制在 64 个字以内",
			Severity: SeverityWarning,
		})
	}
	return issues
}
