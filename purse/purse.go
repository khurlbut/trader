package purse

import (
     "fmt"
     "math"
)

var coins float64
var cashHoldings float64
var targetCashPercentage float64
var tradingFeePercentage float64

func Init(target float64, fee float64) {
     targetCashPercentage = target
     tradingFeePercentage = fee
}

func Fund(funds float64, spot float64) {
     half := 0.5 * funds
     coins = half / spot
     cashHoldings = half
}

func Coins() float64 {
     return coins
}

func CashHoldings() float64 {
     return cashHoldings
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
     target := ValueAt(spot) * targetCashPercentage
     holdings := CashHoldings() 
     adjustment := target - holdings
     fee := tradingFee(adjustment)
     if feeCausesCostOverrun(fee, adjustment, holdings) {
          return adjustment + fee
     }
     return adjustment
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

func feeCausesCostOverrun(fee float64, adjustment float64, holdings float64) bool {
     if adjustment >= holdings {
          return false
     }
     return (fee + math.Abs(adjustment)) > holdings 
}

func addCash(cash float64) float64 {
     cashHoldings = cashHoldings + cash
     return cashHoldings
}

func subCash(cash float64) float64 {
     cashHoldings = cashHoldings - cash
     return cashHoldings
}

func subCoins(c float64) float64 {
     coins = coins - c
     return coins
}

func String(spot float64) string {
     return fmt.Sprintf("Spot: %f\tCashHoldings %f\tCoins: %f\t\tTotal Purse: %f", spot, CashHoldings(), Coins(), ValueAt(spot))
}

func tradingFee(amt float64) float64 {
     return math.Abs(amt) * tradingFeePercentage
}
