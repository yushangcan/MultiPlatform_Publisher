package platform

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

var (
	ErrNilAdapter       = errors.New("platform adapter is nil")
	ErrUnsupported      = errors.New("platform is not supported")
	ErrDuplicateAdapter = errors.New("platform adapter already registered")
)

type Registry struct {
	mu       sync.RWMutex
	adapters map[Platform]Adapter
}

func NewRegistry(adapters ...Adapter) (*Registry, error) {
	registry := &Registry{
		adapters: make(map[Platform]Adapter),
	}

	for _, adapter := range adapters {
		if err := registry.Register(adapter); err != nil {
			return nil, err
		}
	}
	return registry, nil
}

func (registry *Registry) Register(adapter Adapter) error {
	if adapter == nil {
		return ErrNilAdapter
	}

	platform := adapter.Platform()
	if !platform.IsSupported() {
		return fmt.Errorf("%w: %s", ErrUnsupported, platform)
	}

	registry.mu.Lock()
	defer registry.mu.Unlock()

	if _, exists := registry.adapters[platform]; exists {
		return fmt.Errorf("%w: %s", ErrDuplicateAdapter, platform)
	}
	registry.adapters[platform] = adapter
	return nil
}

func (registry *Registry) Get(platform Platform) (Adapter, bool) {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	adapter, ok := registry.adapters[platform]
	return adapter, ok
}

func (registry *Registry) MustGet(platform Platform) (Adapter, error) {
	adapter, ok := registry.Get(platform)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrUnsupported, platform)
	}
	return adapter, nil
}

func (registry *Registry) List() []AdapterInfo {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	platforms := make([]Platform, 0, len(registry.adapters))
	for platform := range registry.adapters {
		platforms = append(platforms, platform)
	}
	sort.Slice(platforms, func(i, j int) bool {
		return platforms[i] < platforms[j]
	})

	result := make([]AdapterInfo, 0, len(platforms))
	for _, platform := range platforms {
		result = append(result, NewAdapterInfo(platform))
	}
	return result
}

func (registry *Registry) Len() int {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	return len(registry.adapters)
}
