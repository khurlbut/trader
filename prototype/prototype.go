package prototype


/* 
     Alternative price sources
     "github.com/khurlbut/trader/price_quotes"
     "github.com/khurlbut/trader/price_quotes/log_based_quotes/cryptodatadownload/price_quotes"
     "github.com/khurlbut/trader/campaign"
*/
import
(
     "fmt"
     "github.com/khurlbut/trader/purse"
     "github.com/khurlbut/trader/price_quotes/log_based_quotes/cryptodatadownload/price_quotes"
     "github.com/khurlbut/trader/cash_percentage_adjuster"
)

var buyTrigger float64
var sellTrigger float64

// var initalCashAmount = 10000.00

func PricingLoop(c *campaign.Campaign) string {
     price_quotes.Init()
     defer price_quotes.Close()

     lastTransctionPrice := price_quotes.CurrentPrice()
     spotPrice := lastTransctionPrice

     cpa := cash_percentage_adjuster.NewCashAdjuster()

     var p *purse.Purse = c.Purse 
     p.SetTargetCashPercentage(cpa.CashPercentageTarget(spotPrice))
     p.Fund(c.InitialCash, spotPrice)

     buyTrigger = d.BuyTrigger
     sellTrigger = d.SellTrigger

     fmt.Printf("%s\n", p.String(spotPrice))

     for price_quotes.HasNextPrice() {
          spotPrice = price_quotes.NextPrice()
          p.SetTargetCashPercentage(cpa.CashPercentageTarget(spotPrice))
          
          if isActionSignaled(spotPrice, lastTransctionPrice) {
               var action string
               cashAdjustmentRequired := p.CashRequiredToAlignWithTarget(spotPrice)

               if isBuy(spotPrice, lastTransctionPrice, cashAdjustmentRequired) {
                    action = "BUY"
                         //
                         // Place BUY order!
                         //
               } else if isSell(spotPrice, lastTransctionPrice, cashAdjustmentRequired) {
                    action = "SELL"
                         //
                         // Place SELL order!
                         //
               }

               if action != "" {
                    p.ReflectOrderFill(cashAdjustmentRequired, spotPrice)
                    lastTransctionPrice = spotPrice

                    fmt.Printf("\t%s\t%s\n", action, p.String(spotPrice))
                    fmt.Printf("\tNew Cash Target: %f\n", cpa.CashPercentageTarget(spotPrice))
               }
          }
     }
     return fmt.Sprintf("%s\n", p.String(spotPrice))
}

func isActionSignaled(spot float64, last float64) bool {
     return isBuySignaled(spot, last) || isSellSignaled(spot, last)
}

func isBuy(spot float64, last float64, buyAmount float64) bool {
     return isBuySignaled(spot, last) && isBuyAdvised(buyAmount)
}

func isBuySignaled(spot float64, last float64) bool {
     if spot < last {
          return delta(spot, last) >= buyTrigger
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
          return delta(spot, last) >= sellTrigger
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