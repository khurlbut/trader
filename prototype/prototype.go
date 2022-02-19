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
var initalFiatAmount = 10000.00

func PricingLoop() string {
     price_quotes.Init()
     defer price_quotes.Close()


     lastTransctionPrice := price_quotes.CurrentPrice()
     spotPrice := lastTransctionPrice
     
     purse.Init((0.5*initalFiatAmount)/spotPrice, 0.5*initalFiatAmount, purseFiatTargetPercent)

     fmt.Printf("Start: %s\n", purse.String(spotPrice))

     for price_quotes.HasNextPrice() {
          spotPrice = price_quotes.NextPrice()
          
          if isActionable(spotPrice, lastTransctionPrice) {
               var action string

               fiatTransactionAmount := purse.FiatRequiredToAlignWithTarget(spotPrice)
               if fiatTransactionAmount == 0 {continue}

               // if isBuy(spotPrice, lastTransctionPrice) {
               if fiatTransactionAmount < 0 {}
                    action = "BUY"
                    // Place buy order for fiatPurchaseAmount worth of crypto
                    purse.AddFiat(fiatTransactionAmount)
                    purse.AddFiat((tradingFee(fiatTransactionAmount)))
                    purse.AddCoins(fiatTransactionAmount * -1 / spotPrice)
               } else {
                    action = "SELL"
                    // Place sell order for cryptoSellAmount of crypto
                    purse.AddFiat(fiatTransactionAmount)
                    purse.AddFiat(tradingFee(fiatTransactionAmount))
                    purse.AddCoins(fiatTransactionAmount / spotPrice)
               }
               lastTransctionPrice = spotPrice
               fmt.Printf("\t" + transactionReport(action, purse.String(spotPrice)))
          }
     }
     return fmt.Sprintf("Final: %s\n", purse.String(spotPrice))
}

func transactionReport(action string, report string) string {
     return fmt.Sprintf("%s\t%s\n", action, report)
}

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

func tradingFee(fiat float64) float64 {
     return fiat * math.Abs(tradingFeePercentage) * -1
}