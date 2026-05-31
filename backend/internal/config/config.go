package config

import (
	"os"
	"strconv"
	"time"
)

const (
	DefaultPort        = "8080"
	DefaultLLMProvider = "rule"
	DefaultLLMModel    = "deepseek-chat"
	DefaultLLMBaseURL  = "https://api.deepseek.com"
	DefaultLLMTimeout  = 30 * time.Second
)

type Config struct {
	Port        string
	LLMProvider string
	LLMAPIKey   string
	LLMModel    string
	LLMBaseURL  string
	LLMTimeout  time.Duration
}

func Load() Config {
	return Config{
		Port:        getEnv("PORT", DefaultPort),
		LLMProvider: getEnv("LLM_PROVIDER", DefaultLLMProvider),
		LLMAPIKey:   os.Getenv("LLM_API_KEY"),
		LLMModel:    getEnv("LLM_MODEL", DefaultLLMModel),
		LLMBaseURL:  getEnv("LLM_BASE_URL", DefaultLLMBaseURL),
		LLMTimeout:  getDurationEnv("LLM_TIMEOUT_SECONDS", DefaultLLMTimeout),
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	seconds, err := strconv.Atoi(value)
	if err != nil || seconds <= 0 {
		return fallback
	}
	return time.Duration(seconds) * time.Second
}
