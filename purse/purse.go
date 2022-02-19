package purse

import "fmt"

var coins float64
var cash float64
var targetCashPercentage float64

func Init(coin float64, csh float64, target float64) {
     coins = coin
     cash = csh
     targetCashPercentage = target
}

func Coins() float64 {
     return coins
}

func Cash() float64 {
     return cash
}

func CoinValue(spot float64) float64 {
     return coins * spot
}

func Value(spot float64) float64 {
     return Cash() + CoinValue(spot)
}

func CashRequiredToAlignWithTarget(spot float64) float64 {
     target := Value(spot) * targetCashPercentage
     return target - Cash()
}

func AddCash(d float64) float64 {
     cash = cash + d
     return cash
}

func AddCoins(c float64) float64 {
     coins = coins + c
     return coins
}

func String(spot float64) string {
     return fmt.Sprintf("Spot: %f\tCash %f\tCoins: %f\t\tTotal Purse: %f", spot, Cash(), Coins(), Value(spot))
}