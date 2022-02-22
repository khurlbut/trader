package price_quotes

import "fmt"

// var spotPriceIndex = 0
// var prices = []float64{5.0, 8.0, 4.0, 2.0, 6.0, 7.0}
// var prices = []float64{1000.0, 1020.0, 1040.4, 1019.592, 998.784, 1000}
// var prices = []float64{5000.0, 7500.00}
// var prices = []float64{7500.0, 5000.00}

type StubQuoteService struct {
     spotPriceIndex int
     prices []float64
}

func NewStubQuoteService() *StubQuoteService {
	return &StubQuoteService{
          spotPriceIndex: 0,
          prices: []float64{7500.0, 5000.00} ,
     }
}

func (qs *StubQuoteService) Open() {
}

func (qs *StubQuoteService) Close() {
}

func (qs *StubQuoteService) HasNextPrice() bool {
     return qs.spotPriceIndex < len(qs.prices)
}

func (qs *StubQuoteService) NextPrice() float64 {
     p := qs.prices[qs.spotPriceIndex]
     qs.spotPriceIndex++
     return p
}

func (qs *StubQuoteService) CurrentPrice() float64 {
     return qs.prices[qs.spotPriceIndex]
}