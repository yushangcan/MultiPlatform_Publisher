package analyzer

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/content"
)

type Handler struct {
	analyzer ContentAnalyzer
}

func NewHandler(analyzer ContentAnalyzer) Handler {
	return Handler{analyzer: analyzer}
}

type AnalyzeRequest struct {
	Input       string `json:"input"`
	ContentType string `json:"content_type,omitempty"`
}

type AnalyzeResponse struct {
	Content content.StructuredContent `json:"content"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func RegisterRoutes(router gin.IRouter, analyzer ContentAnalyzer) {
	handler := NewHandler(analyzer)
	router.POST("/analyze", handler.Analyze)
}

func (handler Handler) Analyze(ctx *gin.Context) {
	var request AnalyzeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid JSON request body"})
		return
	}

	input := content.RawInput{
		Text:        request.Input,
		ContentType: request.ContentType,
	}

	structured, err := handler.analyzer.Analyze(ctx.Request.Context(), input)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, content.ErrRawInputTextRequired) || errors.Is(err, content.ErrRawInputTextTooLong) {
			status = http.StatusBadRequest
		}
		ctx.JSON(status, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, AnalyzeResponse{Content: structured})
}
