package rewrite_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/rewrite"
)

func TestHandlerRewrite(t *testing.T) {
	gin.SetMode(gin.TestMode)

	registry, err := platform.NewRegistry(platform.NewWechatAdapter(), platform.NewZhihuAdapter())
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	router := gin.New()
	rewrite.RegisterRoutes(router, registry)

	body := []byte(`{"content":{"topic":"大学生暑假提升自己","core_points":["学习 Go","完成项目 demo"]},"platforms":["wechat","zhihu"]}`)
	request := httptest.NewRequest(http.MethodPost, "/rewrite", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, response.Code, response.Body.String())
	}

	var payload rewrite.RewriteResponse
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if len(payload.Drafts) != 2 {
		t.Fatalf("expected 2 drafts, got %d", len(payload.Drafts))
	}
}

func TestHandlerRewriteRejectsInvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	registry, err := platform.NewRegistry(platform.NewWechatAdapter())
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	router := gin.New()
	rewrite.RegisterRoutes(router, registry)

	request := httptest.NewRequest(http.MethodPost, "/rewrite", bytes.NewReader([]byte(`{}`)))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}
