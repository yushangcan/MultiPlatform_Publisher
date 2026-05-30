package main

import (
	"log"
	"os"

	"github.com/yushangcan/MultiPlatform_Publisher/backend/internal/router"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	engine := router.New()
	addr := ":" + port

	log.Printf("starting MultiPlatform Publisher API on %s", addr)
	if err := engine.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
