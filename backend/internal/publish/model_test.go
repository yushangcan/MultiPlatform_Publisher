package publish_test

import (
	"testing"
	"time"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/publish"
)

func TestNewPendingTask(t *testing.T) {
	now := time.Date(2026, 5, 30, 16, 30, 0, 0, time.UTC)
	task := publish.NewPendingTask("task-1", []platform.Platform{platform.Wechat, platform.Bilibili}, now)

	if task.ID != "task-1" {
		t.Fatalf("expected task id task-1, got %q", task.ID)
	}
	if task.Status != publish.TaskPending {
		t.Fatalf("expected pending task status, got %s", task.Status)
	}
	if task.Completed() {
		t.Fatal("new pending task should not be completed")
	}
	if len(task.Results) != 2 {
		t.Fatalf("expected 2 platform results, got %d", len(task.Results))
	}
	if task.Results[0].Status != publish.ResultPending {
		t.Fatalf("expected pending result status, got %s", task.Results[0].Status)
	}
}

func TestPublishTaskCompleted(t *testing.T) {
	completedStatuses := []publish.TaskStatus{
		publish.TaskSuccess,
		publish.TaskPartialFailed,
		publish.TaskFailed,
	}

	for _, status := range completedStatuses {
		task := publish.PublishTask{Status: status}
		if !task.Completed() {
			t.Fatalf("expected status %s to be completed", status)
		}
	}
}
