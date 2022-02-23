package price_quotes

import (
)

type QuoteService interface {
	Open()
	Close()
	HasNextPrice() bool
	NextPrice() float64
	CurrentPrice() float64
	// Sleep()
}
