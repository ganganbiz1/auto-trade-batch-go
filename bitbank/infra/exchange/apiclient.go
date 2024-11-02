package exchange

import (
	"context"

	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/config"
	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/domain/entity"
)

type APIClient struct {
	config config.APIConfig
}

func NewAPIClient(config config.APIConfig) *APIClient {
	return &APIClient{config: config}
}

func (c *APIClient) Order(ctx context.Context, order entity.Order) error {
	return nil
}
