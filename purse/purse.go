package purse

import (
     "fmt"
     "math"
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

func ValueAt(spot float64) float64 {
     return CashHoldings() + CoinValue(spot)
}

func CashRequiredToAlignWithTarget(spot float64) float64 {
     target := ValueAt(spot) * targetCashPercentage
     fmt.Printf("CashRequiredToAlignWithTarget: %f\n", target)
     holdings := CashHoldings() 
     adjustment := target - holdings
     fee := tradingFee(adjustment)
     if feeCausesCostOverrun(fee, adjustment, holdings) {
          fmt.Printf("returning adjustment + fee: %f\n", adjustment + fee)
     // os.Exit(0)
          return adjustment + fee
     }
     if feeCausesCoinOverrun(fee, adjustment, spot) {
          fmt.Printf("returning adjustment - fee: %f\n", adjustment - fee)
          // os.Exit(0)
          return adjustment - fee
     }
     // os.Exit(0)
     return adjustment
}

func feeCausesCostOverrun(fee float64, adjustment float64, holdings float64) bool {
     // fmt.Printf("fee: %f, adjustment: %f, holdings: %f\n", fee, adjustment, holdings)
     return (fee + math.Abs(adjustment)) > holdings 
}

func feeCausesCoinOverrun(fee float64, adjustment float64, spot float64) bool {
     return (fee + math.Abs(adjustment)) > CoinValue(spot)
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
     return fmt.Sprintf("Spot: %f\tCashHoldings %f\tCoins: %f\t\tTotal Purse: %f", spot, CashHoldings(), Coins(), ValueAt(spot))
}

func tradingFee(amt float64) float64 {
     return math.Abs(amt) * tradingFeePercentage
}