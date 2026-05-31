package platform

import (
	"fmt"
	"strings"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

func titleOrTopic(structured content.StructuredContent) string {
	if title := strings.TrimSpace(structured.SuggestedTitle); title != "" {
		return title
	}
	return strings.TrimSpace(structured.Topic)
}

func topicOrFallback(structured content.StructuredContent) string {
	if topic := strings.TrimSpace(structured.Topic); topic != "" {
		return strings.TrimRight(topic, " ")
	}
	return "这次分享"
}

func audienceOrFallback(structured content.StructuredContent) string {
	if audience := strings.TrimSpace(structured.Audience); audience != "" {
		return audience
	}
	return "目标读者"
}

func joinTags(values []string, fallback string) []string {
	tags := make([]string, 0, len(values)+1)
	seen := map[string]struct{}{}

	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		tags = append(tags, value)
	}

	if len(tags) == 0 && fallback != "" {
		tags = append(tags, fallback)
	}
	return tags
}

func numberedPoints(points []string) string {
	lines := make([]string, 0, len(points))
	for index, point := range points {
		point = strings.TrimSpace(point)
		if point == "" {
			continue
		}
		lines = append(lines, fmt.Sprintf("%d. %s", index+1, point))
	}
	return strings.Join(lines, "\n")
}
