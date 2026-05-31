package main

import (
	"log"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/config"
	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/router"
)

func main() {
	cfg := config.Load()

	engine := router.New()
	addr := ":" + cfg.Port

	log.Printf("starting MultiPlatform Publisher API on %s", addr)
	if err := engine.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
