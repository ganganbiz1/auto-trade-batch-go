package config

import "github.com/ganganbiz1/auto-trade-batch-go/bitbank/env"

type APIConfig struct {
	APIKey    string
	APISecret string
}

func NewAPIConfig(e env.Env) *APIConfig {
	return &APIConfig{
		APIKey:    e.GetAPIKey(),
		APISecret: e.GetAPISecret(),
	}
}
