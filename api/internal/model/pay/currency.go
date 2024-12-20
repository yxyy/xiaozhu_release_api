package pay

type Currency struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	ExchangeRate float64 `json:"exchange_rate"`
	Code         string  `json:"code"`
}
