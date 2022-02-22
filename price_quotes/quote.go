package price_quotes

import (
)

type QuoteService interface {
	Open()
	Close()
	HasNextPrice() bool
	NextPrice() float64
	CurrentPrice() float64
}

// type Quote struct {
// 	spotPrice	float64
// 	timeStamp time.Time
// }

// func NewQuote() *Quote {
// 	q := Quote{}
// 	return &q
// }