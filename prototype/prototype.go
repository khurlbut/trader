package prototype

import "fmt"

const BuyTrigger = 0.25
const SellTrigger = 0.50
const PurchaseScale = 0.50
const SellScale = 0.50

var cryptoVal = 0
var fiatVal = 1000

func PricingLoop() string {

     // var spotPrice float32  = 3.45
     // var latestTransctionPrice float32 = 22.5
     var spotPrice float32  = 10.0
     var latestTransctionPrice float32 = 5.0
     return fmt.Sprintf("isBuy: %t isSell: %t delta: %f", isBuy(spotPrice, latestTransctionPrice), isSell(spotPrice, latestTransctionPrice), delta(spotPrice, latestTransctionPrice))
}

func isBuy(spot float32, last float32) bool {
     if spot < last {
          return true
     }
     return false
}

func isSell(spot float32, last float32) bool {
     if spot > last {
          return true
     }
     return false
}

func delta(spot float32, last float32) float32 {
     var d = spot - last
     if d < 0 {
          d = d * -1.0
     }
     return d / last
}