package price_quotes

import (
)

type QuoteService interface {
	Open()
	Close()
}

// type Quote struct {
// 	spotPrice	float64
// 	timeStamp time.Time
// }

// func NewQuote() *Quote {
// 	q := Quote{}
// 	return &q
// }