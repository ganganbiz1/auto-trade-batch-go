package entity

type Order struct {
	Pair         CurrencyPair `json:"pair"`
	Amount       string       `json:"amount"`
	Price        string       `json:"price"`
	Side         OrderSide    `json:"side"`
	Type         OrderType    `json:"type"`
	PostOnly     bool         `json:"post_only"`
	TriggerPrice string       `json:"trigger_price"`
}
