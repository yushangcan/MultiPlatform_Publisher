package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status string `json:"status"`
}

func RegisterRoutes(router gin.IRouter) {
	router.GET("/health", Handle)
}

func Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{Status: "ok"})
}
