package entity

type CurrencyPair string

const (
	PairXlmJpy CurrencyPair = "xlm_jpy"
)

func (p CurrencyPair) String() string {
	return string(p)
}

func (t CurrencyPair) IsValid() bool {
	switch t {
	case PairXlmJpy:
		return true
	default:
		return false
	}
}
