package quote_service

import (
	"time"
)

type QuoteService struct {
}

func NewQuoteService() *QuoteService {
	return &QuoteService{}
}

type Quote struct {
	spotPrice	float64
	timeStamp time.Time
}

func NewQuote() *Quote {
	q := Quote{}
	return &q
}