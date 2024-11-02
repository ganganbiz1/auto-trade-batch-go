package entity

type OrderType string

const (
	OrderTypeLimit     OrderType = "limit"
	OrderTypeMarket    OrderType = "market"
	OrderTypeStop      OrderType = "stop"
	OrderTypeStopLimit OrderType = "stop_limit"
)

func (t OrderType) String() string {
	return string(t)
}

func (t OrderType) IsValid() bool {
	switch t {
	case OrderTypeLimit, OrderTypeMarket, OrderTypeStop, OrderTypeStopLimit:
		return true
	default:
		return false
	}
}
