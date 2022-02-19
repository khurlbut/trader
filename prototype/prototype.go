package prototype


/* 
     Alternative price sources
     "github.com/khurlbut/trader/price_quotes"
     "github.com/khurlbut/trader/price_quotes/log_based_quotes/cryptodatadownload/price_quotes"
*/
import
(
     "fmt"
     "math"
     "github.com/khurlbut/trader/price_quotes/log_based_quotes/cryptodatadownload/price_quotes"
)

const BuyTrigger = 0.25
const SellTrigger = 0.05

var purseCoins = 1.00
var purseFiatAmount = 10000.00
const purseFiatTargetPercent = 0.05

const tradingFeePercentage = 0.006

func PricingLoop() string {
     price_quotes.Init()
     defer price_quotes.Close()

     lastTransctionPrice := price_quotes.CurrentPrice()
     spotPrice := lastTransctionPrice
     
     purseCoins = (0.5*purseFiatAmount)/spotPrice
     purseFiatAmount = 0.5*purseFiatAmount

     fmt.Printf("Initial: %s\n", purseValReport(spotPrice))

     for price_quotes.HasNextPrice() {
          spotPrice = price_quotes.NextPrice()
          
          if isActionable(spotPrice, lastTransctionPrice) {
               var action string

               fiatPurseTarget := targetFiatAmount(purseVal(spotPrice))
               fiatTransactionAmount := math.Abs(purseFiatAmount - fiatPurseTarget)

               if isBuy(spotPrice, lastTransctionPrice) {
                    if fiatPurseTarget >= purseFiatAmount {continue}

                    action = "BUY"
                    fmt.Printf("\tfiatPurseTarget (buy): %f\n", fiatPurseTarget)
                    fmt.Printf("\tfiatTransactionAmount (buy): %f\n", fiatTransactionAmount)

                    // Place buy order for fiatPurchaseAmount worth of crypto
                    purseFiatAmount -= (fiatTransactionAmount + tradingFee(fiatTransactionAmount))
                    purseCoins += (fiatTransactionAmount / spotPrice)
               } else if isSell(spotPrice, lastTransctionPrice){
                    if fiatPurseTarget <= purseFiatAmount {continue;}

                    action = "SELL"
                    fmt.Printf("\tfiatPurseTarget (sell): %f\n", fiatPurseTarget)
               
                    // Place sell order for cryptoSellAmount of crypto
                    purseFiatAmount +=  (fiatTransactionAmount - tradingFee(fiatTransactionAmount))
                    purseCoins -= (fiatTransactionAmount / spotPrice)
               }
               lastTransctionPrice = spotPrice
               fmt.Printf("\t" + transactionReport(action, spotPrice))
          }
     }
     return fmt.Sprintf("Final: %s\n", purseValReport(spotPrice))
}

func transactionReport(action string, spot float64) string {
     return fmt.Sprintf("%s\tExecuted: %s\n", action, purseValReport(spot))
}

func targetFiatAmount(purse float64) float64 {
     return purse * purseFiatTargetPercent
}

func isActionable(spot float64, ltp float64) bool {
     return isBuy(spot, ltp) || isSell(spot, ltp)
}

func isBuy(spot float64, last float64) bool {
     if purseFiatAmount > 0 && spot < last {
          return delta(spot, last) >= BuyTrigger
     }
     return false
}

func isSell(spot float64, last float64) bool {
     if purseCoins > 0 && spot > last {
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
     return coinVal(spot) + purseFiatAmount
}

func purseValReport(spot float64) string {
     return fmt.Sprintf("Spot: %f\tFiat %f\tCoin: %f\tTotal: %f", spot, purseFiatAmount, purseCoins, coinVal(spot) + purseFiatAmount)
}

func coinVal(fiatPrice float64) float64 {
     return coinValInFiat(fiatPrice, purseCoins)
}

func coinValInFiat(fiatPrice float64, coinAmount float64) float64 {
     return fiatPrice * coinAmount
}

func tradingFee(fiat float64) float64 {
     return fiat * tradingFeePercentage
}