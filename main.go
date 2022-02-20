package main

import(
	"fmt"
	"github.com/khurlbut/trader/prototype"
	"github.com/khurlbut/trader/purse"
)

const targetCashPercentage = 0.5
const tradingFeePercentage = 0.006

func main() {
	p := purse.NewPurse(targetCashPercentage, tradingFeePercentage)
	fmt.Println(prototype.PricingLoop(p))
}
