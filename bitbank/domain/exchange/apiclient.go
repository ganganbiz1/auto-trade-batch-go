package exchange

import (
	"context"

	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/domain/entity"
)

type APIClient interface {
	GetAsset(ctx context.Context) (entity.Asset, error)
	Order(ctx context.Context, order entity.Order) error
}
