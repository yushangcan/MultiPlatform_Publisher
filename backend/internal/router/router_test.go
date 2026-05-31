package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/health"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/router"
)

func TestNewRegistersHealthRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := router.New()
	request := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	response := httptest.NewRecorder()

	engine.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var payload health.Response
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.Status != "ok" {
		t.Fatalf("expected status ok, got %q", payload.Status)
	}
}

func TestNewRegistersAnalyzeRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := router.New()
	body := []byte(`{"input":"我想写一篇关于实训项目复盘的内容，主要讲项目背景、实现过程和收获。"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/analyze", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	engine.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, response.Code, response.Body.String())
	}

	var payload analyzer.AnalyzeResponse
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if !payload.Content.HasCore() {
		t.Fatal("expected structured content with core information")
	}
}
