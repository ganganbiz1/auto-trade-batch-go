package entity

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

func (t OrderSide) String() string {
	return string(t)
}

func (s OrderSide) IsValid() bool {
	switch s {
	case OrderSideBuy, OrderSideSell:
		return true
	default:
		return false
	}
}
