package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	registry *Registry
}

func NewHandler(registry *Registry) Handler {
	return Handler{registry: registry}
}

type ListPlatformsResponse struct {
	Platforms []AdapterInfo `json:"platforms"`
}

func RegisterRoutes(router gin.IRouter, registry *Registry) {
	handler := NewHandler(registry)
	router.GET("/platforms", handler.List)
}

func (handler Handler) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ListPlatformsResponse{
		Platforms: handler.registry.List(),
	})
}
