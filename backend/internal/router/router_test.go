package router_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
