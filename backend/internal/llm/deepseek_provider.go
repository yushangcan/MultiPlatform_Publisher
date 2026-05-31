package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/config"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

var ErrMissingAPIKey = errors.New("llm api key is required")

type DeepSeekProvider struct {
	apiKey     string
	model      string
	baseURL    string
	httpClient *http.Client
}

func NewDeepSeekProvider(cfg config.Config) (DeepSeekProvider, error) {
	if strings.TrimSpace(cfg.LLMAPIKey) == "" {
		return DeepSeekProvider{}, ErrMissingAPIKey
	}

	timeout := cfg.LLMTimeout
	if timeout <= 0 {
		timeout = config.DefaultLLMTimeout
	}

	return DeepSeekProvider{
		apiKey:  cfg.LLMAPIKey,
		model:   cfg.LLMModel,
		baseURL: strings.TrimRight(cfg.LLMBaseURL, "/"),
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (provider DeepSeekProvider) Name() string {
	return "deepseek"
}

func (provider DeepSeekProvider) Analyze(ctx context.Context, input content.RawInput) (content.StructuredContent, error) {
	if err := input.Validate(); err != nil {
		return content.StructuredContent{}, err
	}

	prompt := fmt.Sprintf("请把下面的模糊内容整理为结构化内容，输出 JSON 字段 topic、audience、content_type、tone、core_points、keywords、suggested_title：\n%s", input.NormalizedText())
	_, err := provider.chat(ctx, prompt)
	if err != nil {
		return content.StructuredContent{}, err
	}

	return content.StructuredContent{}, errors.New("deepseek analyze response parsing is not enabled yet; use rule provider for demo")
}

func (provider DeepSeekProvider) Rewrite(ctx context.Context, structured content.StructuredContent, target platform.Platform) (platform.PlatformDraft, error) {
	prompt := fmt.Sprintf("请把结构化内容改写成适合 %s 平台发布的草稿，输出 JSON 字段 title、body、tags。结构化内容：%+v", target.DisplayName(), structured)
	_, err := provider.chat(ctx, prompt)
	if err != nil {
		return platform.PlatformDraft{}, err
	}

	return platform.PlatformDraft{}, errors.New("deepseek rewrite response parsing is not enabled yet; use rule provider for demo")
}

type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Choices []struct {
		Message chatMessage `json:"message"`
	} `json:"choices"`
}

func (provider DeepSeekProvider) chat(ctx context.Context, prompt string) (string, error) {
	body, err := json.Marshal(chatRequest{
		Model: provider.model,
		Messages: []chatMessage{
			{Role: "system", Content: "你是一个内容结构化和平台改写助手，只输出可解析 JSON。"},
			{Role: "user", Content: prompt},
		},
	})
	if err != nil {
		return "", err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, provider.baseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	request.Header.Set("Authorization", "Bearer "+provider.apiKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := provider.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("llm request failed with status %d", response.StatusCode)
	}

	var payload chatResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return "", err
	}
	if len(payload.Choices) == 0 {
		return "", errors.New("llm response has no choices")
	}
	return payload.Choices[0].Message.Content, nil
}

func newHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: timeout}
}
