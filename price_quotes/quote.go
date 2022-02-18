package price_quotes

// func BuildHello() string {
//      return "Hello, world."
// }

// func BuildHi() string {
//      return "Hi, world."
// }

var spotPriceIndex = 0

// var prices = []float64{5.0, 8.0, 4.0, 2.0, 6.0, 7.0}
var prices = []float64{1000.0, 1020.0, 1040.4, 1019.592, 998.784, 1000}

func hasNextPrice() bool {
     return spotPriceIndex < len(prices)
}

func nextPrice() float64 {
     p := prices[spotPriceIndex]
     spotPriceIndex++
     return p
}

func currentPrice() float64 {
     return prices[spotPriceIndex]
}