package purse

import (
     "fmt"
     "math"
     "os"
)

var coins float64
var cashHoldings float64
var targetCashPercentage float64
var tradingFeePercentage float64

func Init(coin float64, cash float64, target float64, fee float64) {
     coins = coin
     cashHoldings = cash
     targetCashPercentage = target
     tradingFeePercentage = fee
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

func Value(spot float64) float64 {
     return CashHoldings() + CoinValue(spot)
}

func CashRequiredToAlignWithTarget(spot float64) float64 {
     target := Value(spot) * targetCashPercentage
     fee := tradingFee(target)
     holdings := CashHoldings() 
     if feeCausesCostOverrun(fee, target, holdings) {
          return target - holdings + fee
     }
     return target - holdings
}

func feeCausesCostOverrun(fee float64, target float64, holdings float64) bool {
     fmt.Printf("fee: %f, target: %f, holdings: %f\n", fee, target, holdings)
     os.Exit(0)
     return (fee + target + holdings) < 0 
}

func AddCash(cash float64) float64 {
     cashHoldings = cashHoldings + cash
     return cashHoldings
}

func AddCoins(c float64) float64 {
     coins = coins + c
     return coins
}

func String(spot float64) string {
     return fmt.Sprintf("Spot: %f\tCashHoldings %f\tCoins: %f\t\tTotal Purse: %f", spot, CashHoldings(), Coins(), Value(spot))
}

func tradingFee(amt float64) float64 {
     return math.Abs(amt) * tradingFeePercentage * -1
}