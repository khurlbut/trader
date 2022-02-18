package price_quotes

var spotPriceIndex = 0

// var prices = []float64{5.0, 8.0, 4.0, 2.0, 6.0, 7.0}
var prices = []float64{10.00, 10.200, 10.404, 10.19592, 9.98784, 10.00}

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