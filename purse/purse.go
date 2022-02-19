package purse

import "fmt"

var coins float64
var fiat float64
var fiatTargetPercentage float64

func Init(c float64, f float64, ftp float64) {
     fmt.Println("purse.Init")
     coins = c
     fiat = f
     fiatTargetPercentage = ftp
}

func Coins() float64 {
     return coins
}

func Fiat() float64 {
     return fiat
}

func CoinValue(spot float64) float64 {
     return coins * spot
}

func Value(spot float64) float64 {
     return CoinValue(spot) + Fiat()
}

func FiatRequiredToAlignWithTarget(spot float64) float64 {
     target := Value(spot) * fiatTargetPercentage
     return target - Fiat()
}

func AddFiat(f float64) float64 {
     fiat = fiat + f
     return fiat
}

func AddCoins(c float64) float64 {
     coins = coins + c
     return coins
}

func String(spot float64) string {
     return fmt.Sprintf("Spot: %f\tFiat %f\t\tCoin: %f\t\tTotal Purse: %f", spot, Fiat(), Coins(), CoinValue(spot) + Value(spot))
}