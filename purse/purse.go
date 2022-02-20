package purse

import (
     "fmt"
     "math"
)

var coins float64
var cash float64
var targetCashPercentage float64
var tradingFeePercentage float64

const minimum_transaction_amount = 10.00

type Purse struct {

}

func Init(target float64, fee float64) {
     targetCashPercentage = target
     tradingFeePercentage = fee
}

func Fund(funds float64, spot float64) {
     coins = (funds * (1 - targetCashPercentage)) / spot
     cash = funds * targetCashPercentage
}

func Coins() float64 {
     return coins
}

func CashHoldings() float64 {
     return cash
}

func CoinValue(spot float64) float64 {
     return coins * spot
}

func ValueAt(spot float64) float64 {
     return CashHoldings() + CoinValue(spot)
}

/*
     The "adjustment" value is either positive or negative reflecting a BUY or a SELL.
*/
func CashRequiredToAlignWithTarget(spot float64) float64 {
     cashTarget := ValueAt(spot) * targetCashPercentage
     holdings := CashHoldings() 
     adjustment := cashTarget - holdings
     if math.Abs(adjustment) < minimum_transaction_amount {
          return 0
     }
     fee := tradingFee(adjustment)
     if feeCausesCostOverrun(fee, adjustment, holdings) {
          return adjustment + fee
     }
     return adjustment
}

func feeCausesCostOverrun(fee float64, adjustment float64, holdings float64) bool {
     if adjustment >= holdings {
          return false
     }
     return (fee + math.Abs(adjustment)) > holdings 
}

/*
     The "amount" is either positive or negative reflecting a BUY or a SELL.
     A BUY will "add" a negative value to Cash, etc.
 */
func ReflectOrderFill(amount float64, spot float64) {
     addCash(amount)
     subCash(tradingFee(amount))
     subCoins(amount / spot)     
}

func addCash(c float64) {
     cash = cash + c
}

func subCash(c float64) {
     cash = cash - c
}

func subCoins(c float64) {
     coins = coins - c
}

func tradingFee(amt float64) float64 {
     return math.Abs(amt) * tradingFeePercentage
}

func String(spot float64) string {
     return fmt.Sprintf("Spot: %f\tCashHoldings %f\tCoins: %f\t\tTotal Purse: %f", spot, CashHoldings(), Coins(), ValueAt(spot))
}
