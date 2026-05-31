package rewrite

import (
	"context"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

type Service struct {
	registry *platform.Registry
}

func NewService(registry *platform.Registry) Service {
	return Service{registry: registry}
}

func (service Service) Rewrite(ctx context.Context, request RewriteRequest) (RewriteResponse, []platform.ValidationIssue) {
	issues := request.Validate()
	if len(issues) > 0 {
		return RewriteResponse{Content: request.Content, Issues: issues}, issues
	}

	drafts := make([]platform.PlatformDraft, 0, len(request.Platforms))
	for _, target := range request.Platforms {
		adapter, err := service.registry.MustGet(target)
		if err != nil {
			issues = append(issues, platform.ValidationIssue{
				Field:    "platforms",
				Code:     "adapter_not_registered",
				Message:  "platform adapter is not registered",
				Severity: platform.SeverityError,
			})
			continue
		}

		draft, err := adapter.Rewrite(ctx, request.Content)
		if err != nil {
			issues = append(issues, platform.ValidationIssue{
				Field:    "drafts",
				Code:     "rewrite_failed",
				Message:  err.Error(),
				Severity: platform.SeverityError,
			})
			continue
		}
		draft.Warnings = append(draft.Warnings, adapter.Validate(draft)...)
		drafts = append(drafts, draft)
	}

	response := RewriteResponse{
		Content: request.Content,
		Drafts:  drafts,
		Issues:  issues,
	}
	return response, issues
}
