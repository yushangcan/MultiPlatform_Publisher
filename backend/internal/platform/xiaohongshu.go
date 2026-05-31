package platform

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

type XiaohongshuAdapter struct{}

func NewXiaohongshuAdapter() XiaohongshuAdapter {
	return XiaohongshuAdapter{}
}

func (adapter XiaohongshuAdapter) Platform() Platform {
	return Xiaohongshu
}

func (adapter XiaohongshuAdapter) Rewrite(_ context.Context, structured content.StructuredContent) (PlatformDraft, error) {
	topic := topicOrFallback(structured)
	title := titleOrTopic(structured)
	if title == "" {
		title = topic + "｜实用经验分享"
	}

	tags := joinTags(structured.NonEmptyKeywords(), topic)
	body := fmt.Sprintf(
		"%s\n\n适合想要快速行动的人参考。\n\n%s\n\n我的建议是先做最容易开始的一步，不要一开始就追求完美。完成以后再复盘调整，效果会更明显。\n\n%s",
		topic,
		bulletPoints(structured.NonEmptyCorePoints(), "· "),
		hashTags(tags),
	)

	return PlatformDraft{
		Platform:  Xiaohongshu,
		Title:     title,
		Body:      body,
		Tags:      tags,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (adapter XiaohongshuAdapter) Validate(draft PlatformDraft) []ValidationIssue {
	issues := draft.ValidateBasic()
	if len([]rune(draft.Title)) > 30 {
		issues = append(issues, ValidationIssue{
			Field:    "title",
			Code:     "xiaohongshu_title_too_long",
			Message:  "小红书标题建议控制在 30 个字以内，更适合信息流浏览",
			Severity: SeverityWarning,
		})
	}
	if len(draft.Tags) < 3 {
		issues = append(issues, ValidationIssue{
			Field:    "tags",
			Code:     "xiaohongshu_tags_too_few",
			Message:  "小红书建议至少准备 3 个标签",
			Severity: SeverityWarning,
		})
	}
	return issues
}

func hashTags(tags []string) string {
	result := make([]string, 0, len(tags))
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		result = append(result, "#"+strings.ReplaceAll(tag, " ", ""))
	}
	return strings.Join(result, " ")
}
