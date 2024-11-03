package main

import (
	"context"
	"net/http"

	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/config"
	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/env"
	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/infra/exchange"
)

func main() {
	env := env.NewEnv()
	apiConfig := config.NewAPIConfig(env)
	httpClient := &http.Client{}
	apiClient := exchange.NewAPIClient(apiConfig, httpClient)

	apiClient.GetAsset(context.Background())
}
