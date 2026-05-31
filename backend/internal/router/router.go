package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/health"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
)

func New() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	api := engine.Group("/api")
	health.RegisterRoutes(api)
	analyzer.RegisterRoutes(api, analyzer.NewRuleAnalyzer())
	platform.RegisterRoutes(api, defaultPlatformRegistry())

	return engine
}

func defaultPlatformRegistry() *platform.Registry {
	registry, err := platform.NewRegistry(
		platform.NewStaticAdapter(platform.Wechat),
		platform.NewStaticAdapter(platform.Zhihu),
		platform.NewStaticAdapter(platform.Bilibili),
		platform.NewStaticAdapter(platform.Xiaohongshu),
	)
	if err != nil {
		panic(err)
	}
	return registry
}
