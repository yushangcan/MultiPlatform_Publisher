package rewrite

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func RegisterRoutes(router gin.IRouter, registry *platform.Registry) {
	handler := NewHandler(NewService(registry))
	router.POST("/rewrite", handler.Rewrite)
}

func (handler Handler) Rewrite(ctx *gin.Context) {
	var request RewriteRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON request body"})
		return
	}

	response, issues := handler.service.Rewrite(ctx.Request.Context(), request)
	if hasErrorIssue(issues) {
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func hasErrorIssue(issues []platform.ValidationIssue) bool {
	for _, issue := range issues {
		if issue.Severity == platform.SeverityError {
			return true
		}
	}
	return false
}
