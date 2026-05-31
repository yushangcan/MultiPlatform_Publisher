package platform

import (
	"context"
	"fmt"
	"time"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

type ZhihuAdapter struct{}

func NewZhihuAdapter() ZhihuAdapter {
	return ZhihuAdapter{}
}

func (adapter ZhihuAdapter) Platform() Platform {
	return Zhihu
}

func (adapter ZhihuAdapter) Rewrite(_ context.Context, structured content.StructuredContent) (PlatformDraft, error) {
	topic := topicOrFallback(structured)
	title := titleOrTopic(structured)
	if title == "" {
		title = topic + "应该怎么做？"
	}

	body := fmt.Sprintf(
		"这个问题可以先给出一个简短结论：%s 的关键不在于一次性做很多事，而在于找到最值得投入的方向，并持续完成可验证的动作。\n\n我的理解可以拆成几个方面：\n\n%s\n\n如果把它放到实际场景里，我更建议先选一个最重要的切入点，把它做成一个能展示、能复盘的小成果。这样比只停留在计划阶段更有效。\n\n你也可以根据自己的情况调整优先级，重点是让每一步都有产出。",
		topic,
		numberedPoints(structured.NonEmptyCorePoints()),
	)

	return PlatformDraft{
		Platform:  Zhihu,
		Title:     title,
		Body:      body,
		Tags:      joinTags(structured.NonEmptyKeywords(), topic),
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (adapter ZhihuAdapter) Validate(draft PlatformDraft) []ValidationIssue {
	issues := draft.ValidateBasic()
	if len([]rune(draft.Title)) > 80 {
		issues = append(issues, ValidationIssue{
			Field:    "title",
			Code:     "zhihu_title_too_long",
			Message:  "知乎标题建议控制在 80 个字以内",
			Severity: SeverityWarning,
		})
	}
	return issues
}
