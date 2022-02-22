package price_quotes

import "fmt"

var spotPriceIndex = 0

type QuoteService struct {}

func NewStubQuoteService() *QuoteService {
	return &QuoteService{}
}

func Open() {
     fmt.Println("price_quotes.Init")
}

func Init() {
     fmt.Println("price_quotes.Init")
}

func Close() {
     fmt.Println("price_quotes.Close")
}

// var prices = []float64{5.0, 8.0, 4.0, 2.0, 6.0, 7.0}
// var prices = []float64{1000.0, 1020.0, 1040.4, 1019.592, 998.784, 1000}
// var prices = []float64{5000.0, 7500.00}
var prices = []float64{7500.0, 5000.00}

func HasNextPrice() bool {
     return spotPriceIndex < len(prices)
}

func NextPrice() float64 {
     p := prices[spotPriceIndex]
     spotPriceIndex++
     return p
}

func CurrentPrice() float64 {
     return prices[spotPriceIndex]
}