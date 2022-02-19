package prototype


/* 
     Alternative price sources
     "github.com/khurlbut/trader/price_quotes"
     "github.com/khurlbut/trader/purse"
     "github.com/khurlbut/trader/price_quotes/log_based_quotes/cryptodatadownload/price_quotes"
*/
import
(
     "fmt"
     "math"
     "github.com/khurlbut/trader/purse"
     "github.com/khurlbut/trader/price_quotes/log_based_quotes/cryptodatadownload/price_quotes"
)

const BuyTrigger = 0.05
const SellTrigger = 0.05
const purseFiatTargetPercent = 0.0

const tradingFeePercentage = 0.006

// Initial State
var purseFiatAmount = 10000.00

func PricingLoop() string {
     price_quotes.Init()
     defer price_quotes.Close()


     lastTransctionPrice := price_quotes.CurrentPrice()
     spotPrice := lastTransctionPrice
     
     purseFiatAmount = 0.5*purseFiatAmount
     purse.Init((0.5*purseFiatAmount)/spotPrice, 0.5*purseFiatAmount, purseFiatTargetPercent)

     fmt.Printf("Initial: %s\n", purse.String(spotPrice))

     for price_quotes.HasNextPrice() {
          spotPrice = price_quotes.NextPrice()
          
          if isActionable(spotPrice, lastTransctionPrice) {
               var action string

               // fiatPurseTarget := targetFiatAmount(purseVal(spotPrice))
               // fiatTransactionAmount := math.Abs(purseFiatAmount - fiatPurseTarget)
               fiatTransactionAmount := purse.FiatRequiredToAlignWithTarget(spotPrice)

               if isBuy(spotPrice, lastTransctionPrice) {
                    if fiatTransactionAmount >= 0 {continue}

                    action = "BUY"

                    // Place buy order for fiatPurchaseAmount worth of crypto
                    purse.AddFiat(fiatTransactionAmount)
                    purse.AddFiat((tradingFee(fiatTransactionAmount)))
                    purse.AddCoins(fiatTransactionAmount * -1 / spotPrice)
                    // purseFiatAmount += (fiatTransactionAmount + tradingFee(fiatTransactionAmount))
                    // purseCoins += (fiatTransactionAmount / spotPrice)
               } else if isSell(spotPrice, lastTransctionPrice){
                    if fiatTransactionAmount <= 0 {continue}

                    action = "SELL"
               
                    // Place sell order for cryptoSellAmount of crypto
                    purse.AddFiat(fiatTransactionAmount)
                    purse.AddFiat(tradingFee(fiatTransactionAmount))
                    purse.AddCoins(fiatTransactionAmount / spotPrice)
                    // purseFiatAmount +=  (fiatTransactionAmount - tradingFee(fiatTransactionAmount))
                    // purseCoins -= (fiatTransactionAmount / spotPrice)
               }
               lastTransctionPrice = spotPrice
               fmt.Printf("\t" + transactionReport(action, purse.String(spotPrice)))
          }
     }
     return fmt.Sprintf("Final: %s\n", purse.String(spotPrice))
}

func transactionReport(action string, report string) string {
     return fmt.Sprintf("\t%s\n", action, report)
}

// func targetFiatAmount(purse float64) float64 {
//      return purse * purseFiatTargetPercent
// }

func isActionable(spot float64, ltp float64) bool {
     return isBuy(spot, ltp) || isSell(spot, ltp)
}

func isBuy(spot float64, last float64) bool {
     if spot < last {
          return delta(spot, last) >= BuyTrigger
     }
     return false
}

func isSell(spot float64, last float64) bool {
     if spot > last {
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

// func purseVal(spot float64) float64 {
//      return coinVal(spot) + purseFiatAmount
// }

// func purseValReport(spot float64) string {
//      return fmt.Sprintf("Spot: %f\tFiat %f\tCoin: %f\tTotal: %f", spot, purseFiatAmount, purseCoins, coinVal(spot) + purseFiatAmount)
// }

// func coinVal(fiatPrice float64) float64 {
//      return coinValInFiat(fiatPrice, purseCoins)
// }

// func coinValInFiat(fiatPrice float64, coinAmount float64) float64 {
//      return fiatPrice * coinAmount
// }

func tradingFee(fiat float64) float64 {
     return fiat * math.Abs(tradingFeePercentage) * -1
}