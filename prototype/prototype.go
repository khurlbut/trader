package prototype

import "fmt"

const BuyTrigger = 0.25
const SellTrigger = 0.50
const PurchaseScale = 0.50
const SellScale = 0.50

var cryptoVal = 0
var fiatVal = 1000

func PricingLoop() string {
     spotPrice := 23.45
     latestTransctionPrice := 22.5

     return fmt.Sprintf("done %f %f",spotPrice,latestTransctionPrice)
}

func isBuy() bool {
     return true
}