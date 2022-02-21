package quote_service

import (
	"fmt"
	"time"
)

typedef Quote struct {
	spotPrice	float64
	timeStamp time.Time
}

func NewQuote() *Quote {
	q := Quote{}
	return &q
}