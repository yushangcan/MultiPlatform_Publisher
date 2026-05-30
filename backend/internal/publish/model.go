package publish

import (
	"time"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

type TaskStatus string

const (
	TaskPending       TaskStatus = "pending"
	TaskSuccess       TaskStatus = "success"
	TaskPartialFailed TaskStatus = "partial_failed"
	TaskFailed        TaskStatus = "failed"
)

type ResultStatus string

const (
	ResultPending ResultStatus = "pending"
	ResultSuccess ResultStatus = "success"
	ResultFailed  ResultStatus = "failed"
)

type PublishTask struct {
	ID        string                  `json:"id"`
	Status    TaskStatus              `json:"status"`
	Results   []PlatformPublishResult `json:"results"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

func NewPendingTask(id string, targets []platform.Platform, now time.Time) PublishTask {
	results := make([]PlatformPublishResult, 0, len(targets))
	for _, target := range targets {
		results = append(results, PlatformPublishResult{
			Platform: target,
			Status:   ResultPending,
		})
	}

	return PublishTask{
		ID:        id,
		Status:    TaskPending,
		Results:   results,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (task PublishTask) Completed() bool {
	return task.Status == TaskSuccess || task.Status == TaskPartialFailed || task.Status == TaskFailed
}

type PlatformPublishResult struct {
	Platform     platform.Platform `json:"platform"`
	Status       ResultStatus      `json:"status"`
	MockURL      string            `json:"mock_url,omitempty"`
	ErrorMessage string            `json:"error_message,omitempty"`
	PublishedAt  *time.Time        `json:"published_at,omitempty"`
}
