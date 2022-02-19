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
     // fmt.Println("CashRequiredToAlignWithTarget")
     target := Value(spot) * targetCashPercentage
     holdings := CashHoldings() 
     adjustment := target - holdings
     fee := tradingFee(adjustment)
     if feeCausesCostOverrun(fee, adjustment, holdings) {
          fmt.Printf("returning adjustment + fee: %f\n", adjustment + fee)
     // os.Exit(0)
          return adjustment + fee
     }
     // os.Exit(0)
     return adjustment
}

func feeCausesCostOverrun(fee float64, adjustment float64, holdings float64) bool {
     // fmt.Printf("fee: %f, adjustment: %f, holdings: %f\n", fee, adjustment, holdings)
     
     return (fee + math.Abs(adjustment)) > holdings 
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
     return math.Abs(amt) * tradingFeePercentage
}