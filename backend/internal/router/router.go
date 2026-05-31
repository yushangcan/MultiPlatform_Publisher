package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/analyzer"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/health"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/platform"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/rewrite"
)

func New() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	api := engine.Group("/api")
	health.RegisterRoutes(api)
	analyzer.RegisterRoutes(api, analyzer.NewRuleAnalyzer())
	platformRegistry := defaultPlatformRegistry()
	platform.RegisterRoutes(api, platformRegistry)
	rewrite.RegisterRoutes(api, platformRegistry)

	return engine
}

func defaultPlatformRegistry() *platform.Registry {
	registry, err := platform.NewRegistry(
		platform.NewWechatAdapter(),
		platform.NewZhihuAdapter(),
		platform.NewBilibiliAdapter(),
		platform.NewXiaohongshuAdapter(),
	)
	if err != nil {
		panic(err)
	}
	return registry
}
