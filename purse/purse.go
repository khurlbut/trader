package purse

import "fmt"

var coins float64
var cashHoldings float64
var targetcashHoldingsPercentage float64

func Init(coin float64, cash float64, target float64, fee float54) {
     coins = coin
     cashHoldings = cash
     targetcashHoldingsPercentage = target
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
     return cashHoldings() + CoinValue(spot)
}

func CasRequiredToAlignWithTarget(spot float64) float64 {
     target := Value(spot) * targetcashPercentage
     return target - CashHoldings()
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