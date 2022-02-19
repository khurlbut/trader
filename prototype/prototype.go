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
     "github.com/khurlbut/trader/purse"
     "github.com/khurlbut/trader/price_quotes"
)

const BuyTrigger = 0.05
const SellTrigger = 0.05
const targetCashPercentage = 1.0

const tradingFeePercentage = 0.006

// Initial State
var initalCashAmount = 10000.00

func PricingLoop() string {
     price_quotes.Init()
     defer price_quotes.Close()

     lastTransctionPrice := price_quotes.CurrentPrice()
     spotPrice := lastTransctionPrice
     
     purse.Init((0.5*initalCashAmount)/spotPrice, 0.5*initalCashAmount, targetCashPercentage, tradingFeePercentage)

     fmt.Printf("%s\n", purse.String(spotPrice))

     for price_quotes.HasNextPrice() {
          spotPrice = price_quotes.NextPrice()
          
          if isActionSignaled(spotPrice, lastTransctionPrice) {
               var action string

               cashAdjustmentRequired := purse.CashRequiredToAlignWithTarget(spotPrice)

               if isBuy(spotPrice, lastTransctionPrice, cashAdjustmentRequired) {
                    action = "BUY"
                    // Place buy order for fiatPurchaseAmount worth of crypto
                    purse.AddCash(cashAdjustmentRequired)
                    purse.AddCash((tradingFee(cashAdjustmentRequired)))
                    purse.AddCoins(cashAdjustmentRequired * -1 / spotPrice)
               } else if isSell(spotPrice, lastTransctionPrice, cashAdjustmentRequired) {
                    action = "SELL"
                    // Place sell order for cryptoSellAmount of crypto
                    purse.AddCash(cashAdjustmentRequired)
                    purse.AddCash(tradingFee(cashAdjustmentRequired))
                    purse.AddCoins(cashAdjustmentRequired * -1 / spotPrice)
               }

               if action != "" {
                    lastTransctionPrice = spotPrice
                    fmt.Printf("\t" + transactionReport(action, purse.String(spotPrice)))
               }
          }
     }
     return fmt.Sprintf("%s\n", purse.String(spotPrice))
}

func transactionReport(action string, report string) string {
     return fmt.Sprintf("%s\t%s\n", action, report)
}

func isActionSignaled(spot float64, last float64) bool {
     return isBuySignaled(spot, last) || isSellSignaled(spot, last)
}

func isBuy(spot float64, last float64, buyAmount float64) bool {
     return isBuySignaled(spot, last) && isBuyAdvised(buyAmount)
}

func isBuySignaled(spot float64, last float64) bool {
     if spot < last {
          return delta(spot, last) >= BuyTrigger
     }
     return false
}

func isBuyAdvised(amt float64) bool {
     return amt < 0
}

func isSell(spot float64, last float64, sellAmount float64) bool {
     return isSellSignaled(spot, last) && isSellAdvised(sellAmount)
}

func isSellSignaled(spot float64, last float64) bool {
     if spot > last {
          return delta(spot, last) >= SellTrigger
     }
     return false
}

func isSellAdvised(amt float64) bool {
     return amt > 0
}

func delta(spot float64, last float64) float64 {
     var d = spot - last
     if d < 0 {
          d = d * -1.0
     }
     return d / last
}

func tradingFee(amt float64) float64 {
     return math.Abs(amt) * tradingFeePercentage * -1
}