package entity

type OrderStatus string

const (
	OrderStatusInactive                OrderStatus = "INACTIVE"
	OrderStatusUnfilled                OrderStatus = "UNFILLED"
	OrderStatusPartiallyFilled         OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFullyFilled             OrderStatus = "FULLY_FILLED"
	OrderStatusCanceledUnfilled        OrderStatus = "CANCELED_UNFILLED"
	OrderStatusCanceledPartiallyFilled OrderStatus = "CANCELED_PARTIALLY_FILLED"
)

func (s OrderStatus) String() string {
	return string(s)
}

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusInactive,
		OrderStatusUnfilled,
		OrderStatusPartiallyFilled,
		OrderStatusFullyFilled,
		OrderStatusCanceledUnfilled,
		OrderStatusCanceledPartiallyFilled:
		return true
	default:
		return false
	}
}

func (s OrderStatus) Jpy() string {
	switch s {
	case OrderStatusInactive:
		return "非アクティブ"
	case OrderStatusUnfilled:
		return "注文中"
	case OrderStatusPartiallyFilled:
		return "注文中(一部約定)"
	case OrderStatusFullyFilled:
		return "約定済み"
	case OrderStatusCanceledUnfilled:
		return "取消済"
	case OrderStatusCanceledPartiallyFilled:
		return "取消済(一部約定)"
	default:
		return ""
	}
}
