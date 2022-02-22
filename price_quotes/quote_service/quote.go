package quote_service

import (
	"time"
)

type Quote struct {
	spotPrice	float64
	timeStamp time.Time
}

func NewQuote() *Quote {
	q := Quote{}
	return &q
}