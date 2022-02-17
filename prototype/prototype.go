package prototype

import "fmt"

const BuyTrigger = 0.25
const SellTrigger = 0.50
const PurchaseScale = 0.50
const SellScale = 0.50

var cryptoVal = 50.00
var fiatVal = 1000.00

func PricingLoop() string {

     // var spotPrice float64  = 3.45
     // var lastTransctionPrice float64 = 22.5
     // var spotPrice float64  = 5.0
     // var lastTransctionPrice float64 = 10.0
     var spotPrice float64  = 10.0
     var lastTransctionPrice float64 = 5.0
     buy := isBuy(spotPrice, lastTransctionPrice)
     sell := isSell(spotPrice, lastTransctionPrice)
     d := delta(spotPrice, lastTransctionPrice)

     fmt.Sprintf("spot: %f last: %f isBuy: %t isSell: %t delta: %f\n", spotPrice, lastTransctionPrice, buy, sell, d)

     if isBuy(spotPrice, lastTransctionPrice) {
          var fiatPurchaseAmount float64 = PurchaseScale * fiatVal
          // Place buy order for fiatPurchaseAmount worth of crypto
          fiatVal -= fiatPurchaseAmount
          cryptoVal += (fiatPurchaseAmount / spotPrice)
          lastTransctionPrice  = spotPrice
          fmt.Printf("BUY Executed: fiatVal: %f cryptoVal: %f\n", fiatVal, cryptoVal)
     } else if isSell(spotPrice, lastTransctionPrice){
          cryptoSellAmount := SellScale * cryptoVal
          // Place sell order for cryptoSellAmount of crypto
          cryptoVal += cryptoSellAmount
          fiatVal -= (cryptoSellAmount * spotPrice)
          lastTransctionPrice = spotPrice
          fmt.Printf("SELL Executed: fiatVal: %f cryptoVal: %f\n", fiatVal, cryptoVal)
     }

     return "done"
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