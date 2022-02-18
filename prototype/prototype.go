package prototype


/* 
     Alternative price sources
     "github.com/khurlbut/trader/price_quotes"
     "github.com/khurlbut/trader/price_quotes/log_based_quotes/cryptodatadownload/price_quotes"
*/
import
(
     "fmt"
     "github.com/khurlbut/trader/price_quotes/log_based_quotes/cryptodatadownload/price_quotes"
)

const BuyTrigger = 0.05
const SellTrigger = 0.05
const fiatPercentageTarget = 0.25
// const PurchaseScale = 0.250
// const SellScale = 0.40
const tradingFeePercentage = 0.006

var coinCount = 1.00
var fiatVal = 10000.00

func PricingLoop() string {
     price_quotes.Init()
     defer price_quotes.Close()

     lastTransctionPrice := price_quotes.CurrentPrice()
     spotPrice := lastTransctionPrice
     
     coinCount = (0.5*fiatVal)/spotPrice
     fiatVal = 0.5*fiatVal

     fmt.Printf("Initial: %s\n", purseValReport(spotPrice))

     for price_quotes.HasNextPrice() {
          spotPrice = price_quotes.NextPrice()
          
          // buy := isBuy(spotPrice, lastTransctionPrice)
          // sell := isSell(spotPrice, lastTransctionPrice)
          // d := delta(spotPrice, lastTransctionPrice)
          // fmt.Printf("spot: %f last: %f isBuy: %t isSell: %t delta: %f\n", spotPrice, lastTransctionPrice, buy, sell, d)

          if isBuy(spotPrice, lastTransctionPrice) {
               // fiatTransactionVal := PurchaseScale * fiatVal
               fiatTransactionVal := targetFiatAmount(purseVal(spotPrice))
               fmt.Printf("\tfiatTranscationVal (buy): %f\n", fiatTransactionVal)

               // Place buy order for fiatPurchaseAmount worth of crypto

               fiatVal -= (fiatTransactionVal + tradingFee(fiatTransactionVal))
               coinCount += (fiatTransactionVal / spotPrice)
               lastTransctionPrice  = spotPrice

               fmt.Printf("\t" + transactionReport("BUY", spotPrice))
          } else if isSell(spotPrice, lastTransctionPrice){
               // cryptoSellAmount := SellScale * coinCount
               // fiatTransactionVal := coinValInFiat(spotPrice, cryptoSellAmount)  
               fiatTransactionVal := targetFiatAmount(purseVal(spotPrice))
               fmt.Printf("\tfiatTranscationVal (sell): %f\n", fiatTransactionVal)
          
               // Place sell order for cryptoSellAmount of crypto

               fiatVal +=  (fiatTransactionVal - tradingFee(fiatTransactionVal))
               // coinCount -= cryptoSellAmount
               coinCount -= (fiatTransactionVal / spotPrice)
               lastTransctionPrice = spotPrice

               fmt.Printf("\t" + transactionReport("SELL", spotPrice))
          }
          // fmt.Printf("New purse Value: %f\n", purseVal(spotPrice))

     }
     return fmt.Sprintf("Final: %s\n", purseValReport(spotPrice))
}

func transactionReport(action string, spot float64) string {
     return fmt.Sprintf("%s\tExecuted: %s\n", action, purseValReport(spot))
}

func targetFiatAmount(purse float64) float64 {
     return purse * fiatPercentageTarget
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

func purseVal(spot float64) float64 {
     return coinVal(spot) + fiatVal
}

func purseValReport(spot float64) string {
     return fmt.Sprintf("Spot: %f\tFiat %f\tCoin: %f\tTotal: %f", spot, fiatVal, coinCount, coinVal(spot) + fiatVal)
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