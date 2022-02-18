package prototype

import
(
     "fmt"
     "github.com/khurlbut/trader/price_quotes"
)

const BuyTrigger = 0.04
const SellTrigger = 0.02
const PurchaseScale = 0.250
const SellScale = 0.40
const tradingFeePercentage = 0.006

var coinCount = 1.00
var fiatVal = 0.00

func PricingLoop() string {
     lastTransctionPrice := price_quotes.CurrentPrice()
     spotPrice := lastTransctionPrice
     
     fmt.Printf("Initial Wallet Value: %f\n", walletVal(spotPrice))

     for price_quotes.HasNextPrice() {
          spotPrice = price_quotes.NextPrice()
          buy := isBuy(spotPrice, lastTransctionPrice)
          sell := isSell(spotPrice, lastTransctionPrice)
          d := delta(spotPrice, lastTransctionPrice)

          fmt.Printf("spot: %f last: %f isBuy: %t isSell: %t delta: %f\n", spotPrice, lastTransctionPrice, buy, sell, d)

          if isBuy(spotPrice, lastTransctionPrice) {
               fiatTransactionVal := PurchaseScale * fiatVal

               // Place buy order for fiatPurchaseAmount worth of crypto

               fiatVal -= (fiatTransactionVal + tradingFee(fiatTransactionVal))
               coinCount += (fiatTransactionVal / spotPrice)
               lastTransctionPrice  = spotPrice

               fmt.Printf("\tBUY Executed: fiatVal: %f coinCount: %f\n", fiatVal, coinCount)
          } else if isSell(spotPrice, lastTransctionPrice){
               cryptoSellAmount := SellScale * coinCount
               fiatTransactionVal := coinValInFiat(spotPrice, cryptoSellAmount)  
          
               // Place sell order for cryptoSellAmount of crypto

               fiatVal +=  (fiatTransactionVal - tradingFee(fiatTransactionVal))
               coinCount -= cryptoSellAmount
               lastTransctionPrice = spotPrice

               fmt.Printf("\tSELL Executed: fiatVal: %f coinCount: %f\n", fiatVal, coinCount)
          }
          fmt.Printf("New Wallet Value: %f\n", walletVal(spotPrice))

     }
     return fmt.Sprintf("Final Wallet Value: %f\n", walletVal(spotPrice))
}

func isBuy(spot float64, last float64) bool {
     if fiatVal > 0 && spot < last {
          return delta(spot, last) >= BuyTrigger
     }
     return false
}

func isSell(spot float64, last float64) bool {
     if coinCount > 0 && spot > last {
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

func walletVal(spot float64) float64 {
     return coinVal(spot) + fiatVal
}

func coinVal(fiatPrice float64) float64 {
     return coinValInFiat(fiatPrice, coinCount)
}

func coinValInFiat(fiatPrice float64, coinAmount float64) float64 {
     return fiatPrice * coinAmount
}

func tradingFee(fiat float64) float64 {
     return fiat * tradingFeePercentage
}