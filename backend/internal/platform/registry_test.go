package platform_test

import (
	"errors"
	"testing"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func TestRegistryRegisterAndGet(t *testing.T) {
	adapter := platform.NewStaticAdapter(platform.Wechat)
	registry, err := platform.NewRegistry(adapter)
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	got, ok := registry.Get(platform.Wechat)
	if !ok {
		t.Fatal("expected wechat adapter to be registered")
	}
	if got.Platform() != platform.Wechat {
		t.Fatalf("expected wechat adapter, got %s", got.Platform())
	}
}

func TestRegistryRejectsDuplicateAdapter(t *testing.T) {
	registry, err := platform.NewRegistry(platform.NewStaticAdapter(platform.Wechat))
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	err = registry.Register(platform.NewStaticAdapter(platform.Wechat))
	if !errors.Is(err, platform.ErrDuplicateAdapter) {
		t.Fatalf("expected ErrDuplicateAdapter, got %v", err)
	}
}

func TestRegistryRejectsUnsupportedPlatform(t *testing.T) {
	_, err := platform.NewRegistry(platform.NewStaticAdapter(platform.Platform("unknown")))
	if !errors.Is(err, platform.ErrUnsupported) {
		t.Fatalf("expected ErrUnsupported, got %v", err)
	}
}

func TestRegistryList(t *testing.T) {
	registry, err := platform.NewRegistry(
		platform.NewStaticAdapter(platform.Zhihu),
		platform.NewStaticAdapter(platform.Wechat),
	)
	if err != nil {
		t.Fatalf("NewRegistry returned error: %v", err)
	}

	items := registry.List()
	if len(items) != 2 {
		t.Fatalf("expected 2 platforms, got %d", len(items))
	}
	if items[0].DisplayName == "" || items[1].DisplayName == "" {
		t.Fatal("expected display names")
	}
}
