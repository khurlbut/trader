package prototype

import "fmt"

const BuyTrigger = 0.25
const SellTrigger = 0.50
const PurchaseScale = 0.50
const SellScale = 0.50

var cryptoVal = 50.00
var fiatVal = 1000.00

var spotPriceIndex = 0

// var prices = []float64{5.0, 8.0, 4.0, 2.0, 6.0, 7.0}
var prices = []float64{5.0, 8.0, 4.0, 2.0, 6.0, 7.0}

func PricingLoop() string {
     var spotPrice float64
     lastTransctionPrice := 7.0
     
     fmt.Printf("Initial Wallet Value: %f\n", walletVal(lastTransctionPrice))

     for hasNextPrice() {
          spotPrice = nextPrice()
          // fmt.Println(sp)
          buy := isBuy(spotPrice, lastTransctionPrice)
          sell := isSell(spotPrice, lastTransctionPrice)
          d := delta(spotPrice, lastTransctionPrice)

          fmt.Printf("spot: %f last: %f isBuy: %t isSell: %t delta: %f\n", spotPrice, lastTransctionPrice, buy, sell, d)

          if isBuy(spotPrice, lastTransctionPrice) {
               var fiatPurchaseAmount float64 = PurchaseScale * fiatVal
               // Place buy order for fiatPurchaseAmount worth of crypto
               fiatVal -= fiatPurchaseAmount
               cryptoVal += (fiatPurchaseAmount / spotPrice)
               lastTransctionPrice  = spotPrice
               fmt.Printf("\tBUY Executed: fiatVal: %f cryptoVal: %f\n", fiatVal, cryptoVal)
          } else if isSell(spotPrice, lastTransctionPrice){
               cryptoSellAmount := SellScale * cryptoVal
               // Place sell order for cryptoSellAmount of crypto
               fiatVal += (cryptoSellAmount * spotPrice)
               cryptoVal -= cryptoSellAmount
               lastTransctionPrice = spotPrice
               fmt.Printf("\tSELL Executed: fiatVal: %f cryptoVal: %f\n", fiatVal, cryptoVal)
          }
          fmt.Printf("New Wallet Value: %f\n", walletVal(lastTransctionPrice))

     }
     return fmt.Sprintf("Final Wallet Value: %f\n", walletVal(lastTransctionPrice))
}

func isBuy(spot float64, last float64) bool {
     if fiatVal > 0 && spot < last {
          return delta(spot, last) >= BuyTrigger
     }
     return false
}

func isSell(spot float64, last float64) bool {
     if cryptoVal > 0 && spot > last {
          return delta(spot, last) >= SellTrigger
     }
     return false
}

func delta(spot float64, last float64) float64 {
     var d = spot - last
     if d < 0 {
          d = d * -1.0
     }
     return d / last
}

func hasNextPrice() bool {
     return spotPriceIndex < len(prices)
}

func nextPrice() float64 {
     p := prices[spotPriceIndex]
     spotPriceIndex++
     return p
}

func walletVal(spot float64) float64 {
     f := spot * cryptoVal
     return f + fiatVal
}