package price_quotes

import (
  "github.com/khurlbut/trader/price_quotes/stub/price_quotes"
)

type QuoteService interface {
	Open()
	Close()
}

func NewStubQuoteService() *QuoteService {
	return &QuoteService{}
}

// type Quote struct {
// 	spotPrice	float64
// 	timeStamp time.Time
// }

// func NewQuote() *Quote {
// 	q := Quote{}
// 	return &q
// }