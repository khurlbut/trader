package prototype

import "fmt"

const BuyTrigger = 0.25
const SellTrigger = 0.50
const PurchaseScale = 0.50
const SellScale = 0.50

var cryptoVal = 0
var fiatVal = 1000

func PricingLoop() string {
     var spotPrice float32  = 3.45
     var latestTransctionPrice float32 = 22.5

     return fmt.Sprintf("isBuy: %t isSell: %t delta: %f", isBuy(spotPrice, latestTransctionPrice), isSell(spotPrice, latestTransctionPrice), delta(spot, latestTransctionPrice))
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
     if spot < last {
          d = last - spot
     }
     return d / last
}