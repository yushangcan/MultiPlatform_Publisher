package platform_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestHandlerList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	registry, err := platform.NewRegistry(
		platform.NewStaticAdapter(platform.Wechat),
		platform.NewStaticAdapter(platform.Zhihu),
	)
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	router := gin.New()
	platform.RegisterRoutes(router, registry)

	request := httptest.NewRequest(http.MethodGet, "/platforms", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var payload platform.ListPlatformsResponse
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if len(payload.Platforms) != 2 {
		t.Fatalf("expected 2 platforms, got %d", len(payload.Platforms))
	}
}
