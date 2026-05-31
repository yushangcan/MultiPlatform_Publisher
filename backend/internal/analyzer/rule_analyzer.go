package analyzer

import (
	"context"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

const (
	defaultAudience    = "内容创作者关注的目标读者"
	defaultContentType = "经验分享"
	defaultTone        = "清晰、实用"
	maxPointCount      = 6
	maxKeywordCount    = 8
)

var splitPattern = regexp.MustCompile(`[，,。.!！?？；;\n\r]+`)

type RuleAnalyzer struct{}

func NewRuleAnalyzer() RuleAnalyzer {
	return RuleAnalyzer{}
}

func (analyzer RuleAnalyzer) Analyze(_ context.Context, input content.RawInput) (content.StructuredContent, error) {
	if err := input.Validate(); err != nil {
		return content.StructuredContent{}, err
	}

	text := input.NormalizedText()
	points := extractCorePoints(text)
	topic := inferTopic(text, points)
	now := time.Now().UTC()

	return content.StructuredContent{
		SourceID:       input.ID,
		Topic:          topic,
		Audience:       inferAudience(text),
		ContentType:    inferContentType(text, input.ContentType),
		Tone:           inferTone(text),
		CorePoints:     points,
		Keywords:       extractKeywords(text, topic, points),
		SuggestedTitle: suggestTitle(topic),
		CreatedAt:      now,
	}, nil
}

func inferTopic(text string, points []string) string {
	candidates := []string{"关于", "围绕", "主题是", "写一篇", "写一个", "做一期"}
	for _, marker := range candidates {
		if index := strings.Index(text, marker); index >= 0 {
			fragment := firstSegment(text[index+len(marker):])
			fragment = cleanTopic(fragment)
			if fragment != "" {
				return limitRunes(fragment, 28)
			}
		}
	}

	if len(points) > 0 {
		return limitRunes(points[0], 28)
	}
	return limitRunes(trimSentence(text), 28)
}

func inferAudience(text string) string {
	audienceMarkers := []struct {
		marker   string
		audience string
	}{
		{"大学生", "大学生"},
		{"实习生", "实习生和求职学生"},
		{"程序员", "程序员"},
		{"开发者", "开发者"},
		{"创作者", "内容创作者"},
		{"新手", "新手用户"},
		{"学生", "学生群体"},
	}

	for _, item := range audienceMarkers {
		if strings.Contains(text, item.marker) {
			return item.audience
		}
	}
	return defaultAudience
}

func inferContentType(text string, requested string) string {
	if requested = strings.TrimSpace(requested); requested != "" {
		return requested
	}

	contentTypeMarkers := []struct {
		marker      string
		contentType string
	}{
		{"复盘", "复盘总结"},
		{"教程", "教程"},
		{"笔记", "学习笔记"},
		{"产品", "产品介绍"},
		{"测评", "测评推荐"},
		{"经验", "经验分享"},
		{"提升", "经验分享"},
		{"项目介绍", "项目介绍"},
		{"项目", "项目经验"},
	}

	for _, item := range contentTypeMarkers {
		if strings.Contains(text, item.marker) {
			return item.contentType
		}
	}
	return defaultContentType
}

func inferTone(text string) string {
	toneMarkers := []struct {
		marker string
		tone   string
	}{
		{"轻松", "轻松、实用"},
		{"正式", "正式、清晰"},
		{"专业", "专业、严谨"},
		{"口语", "口语化、自然"},
		{"活泼", "活泼、有感染力"},
	}

	for _, item := range toneMarkers {
		if strings.Contains(text, item.marker) {
			return item.tone
		}
	}
	return defaultTone
}

func extractCorePoints(text string) []string {
	segments := splitPattern.Split(text, -1)
	points := make([]string, 0, maxPointCount)
	seen := map[string]struct{}{}

	for _, segment := range segments {
		point := trimSentence(segment)
		if point == "" || utf8.RuneCountInString(point) < 3 {
			continue
		}
		point = removeLeadPhrase(point)
		if point == "" {
			continue
		}
		key := strings.ToLower(point)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		points = append(points, limitRunes(point, 42))
		if len(points) >= maxPointCount {
			break
		}
	}

	if len(points) == 0 {
		points = append(points, limitRunes(text, 42))
	}
	return points
}

func extractKeywords(text string, topic string, points []string) []string {
	candidates := make([]string, 0, maxKeywordCount)
	candidates = append(candidates, knownKeywords(text)...)
	candidates = append(candidates, tokenize(topic)...)
	for _, point := range points {
		candidates = append(candidates, tokenize(point)...)
	}
	candidates = append(candidates, tokenize(text)...)

	keywords := make([]string, 0, maxKeywordCount)
	seen := map[string]struct{}{}
	for _, candidate := range candidates {
		candidate = strings.TrimSpace(candidate)
		if candidate == "" || isStopWord(candidate) {
			continue
		}
		if utf8.RuneCountInString(candidate) > 16 {
			continue
		}
		key := strings.ToLower(candidate)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		keywords = append(keywords, candidate)
		if len(keywords) >= maxKeywordCount {
			break
		}
	}
	return keywords
}

func suggestTitle(topic string) string {
	topic = strings.TrimSpace(topic)
	if topic == "" {
		return "把想法整理成可发布内容"
	}
	return topic + "：一篇适合多平台发布的内容"
}

func trimSentence(value string) string {
	value = strings.TrimSpace(value)
	value = strings.Trim(value, "：:、-— ")
	return value
}

func firstSegment(value string) string {
	segments := splitPattern.Split(value, 2)
	if len(segments) == 0 {
		return trimSentence(value)
	}
	return trimSentence(segments[0])
}

func cleanTopic(value string) string {
	value = trimSentence(value)
	suffixes := []string{"的内容", "的文章", "这件事", "这个话题"}
	for _, suffix := range suffixes {
		value = strings.TrimSuffix(value, suffix)
	}
	return trimSentence(value)
}

func removeLeadPhrase(value string) string {
	prefixes := []string{"我想", "想要", "主要讲", "主要包括", "内容是", "讲", "写"}
	for _, prefix := range prefixes {
		value = strings.TrimSpace(value)
		if strings.HasPrefix(value, prefix) {
			value = strings.TrimSpace(strings.TrimPrefix(value, prefix))
		}
	}
	return trimSentence(value)
}

func limitRunes(value string, limit int) string {
	runes := []rune(strings.TrimSpace(value))
	if len(runes) <= limit {
		return string(runes)
	}
	return string(runes[:limit])
}

func tokenize(text string) []string {
	text = strings.NewReplacer(
		"、", " ",
		"，", " ",
		"。", " ",
		"；", " ",
		"：", " ",
		",", " ",
		".", " ",
		";", " ",
		":", " ",
		"\n", " ",
		"\r", " ",
	).Replace(text)
	return strings.Fields(text)
}

func knownKeywords(text string) []string {
	ordered := []string{
		"大学生",
		"暑假",
		"自我提升",
		"学习 Go",
		"Go",
		"项目",
		"简历",
		"运动",
		"实训",
		"复盘",
		"创作者",
		"多平台发布",
	}

	keywords := make([]string, 0, maxKeywordCount)
	for _, keyword := range ordered {
		if strings.Contains(text, keyword) {
			keywords = append(keywords, keyword)
		}
	}
	return keywords
}

func isStopWord(value string) bool {
	stopWords := map[string]struct{}{
		"我想": {},
		"一个": {},
		"一篇": {},
		"关于": {},
		"如何": {},
		"这个": {},
		"这些": {},
		"主要": {},
		"适合": {},
		"内容": {},
		"希望": {},
	}
	_, ok := stopWords[value]
	return ok
}
