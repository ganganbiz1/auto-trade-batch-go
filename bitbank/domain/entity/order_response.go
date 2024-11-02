package entity

type OrderResponse struct {
	OrderID         int          `json:"order_id"`
	Pair            CurrencyPair `json:"pair"`
	Side            OrderSide    `json:"side"`
	PositionSide    *string      `json:"position_side"`
	Type            OrderType    `json:"type"`
	StartAmount     string       `json:"start_amount"`
	RemainingAmount string       `json:"remaining_amount"`
	ExecutedAmount  string       `json:"executed_amount"`
	Price           *string      `json:"price"`
	PostOnly        bool         `json:"post_only"`
	UserCancelable  bool         `json:"user_cancelable"`
	AveragePrice    string       `json:"average_price"`
	OrderedAt       int          `json:"ordered_at"`
	ExpireAt        *int         `json:"expire_at"`
	TriggeredAt     *int         `json:"triggered_at"`
	TriggerPrice    string       `json:"trigger_price"`
	Status          OrderStatus  `json:"status"`
}
