package analyzer_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
)

func TestHandlerAnalyze(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	analyzer.RegisterRoutes(router, analyzer.NewRuleAnalyzer())

	body := []byte(`{"input":"我想写一篇关于大学生暑假提升的内容，主要讲学习 Go、做项目和准备简历。","content_type":"经验分享"}`)
	request := httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, response.Code, response.Body.String())
	}

	var payload analyzer.AnalyzeResponse
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.Content.Topic == "" {
		t.Fatal("expected structured topic")
	}
	if len(payload.Content.CorePoints) == 0 {
		t.Fatal("expected core points")
	}
}

func TestHandlerAnalyzeRejectsBlankInput(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	analyzer.RegisterRoutes(router, analyzer.NewRuleAnalyzer())

	body := []byte(`{"input":"   "}`)
	request := httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}
