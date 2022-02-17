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
     isBuy := isBuy(spotPrice, latestTransctionPrice)
     isSell := isSell(spotPrice, latestTransctionPrice)
     delta := delta(spotPrice, latestTransctionPrice)
     return fmt.Sprintf("spot: %f last: %f isBuy: %t isSell: %t delta: %f", spotPrice, latestTransctionPrice, isBuy, isSell, delta)
     // var spotPrice float32  = 5.0
     // var latestTransctionPrice float32 = .0
     // var spotPrice float32  = 10.0
     // var latestTransctionPrice float32 = 5.0
}

func isBuy(spot float32, last float32) bool {
     if spot < last {
          return delta(spot, last) >= BuyTrigger
     }
     return false
}

func isSell(spot float32, last float32) bool {
     if spot > last {
          return delta(spot, last) >= SellTrigger
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