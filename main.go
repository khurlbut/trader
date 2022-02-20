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
	// p.FundPurse(10000, 1000)
	// fmt.Printf("\n%s", p.Properties())
	// fmt.Printf("Holdings: %s\n", p.Holdings(1000))
	fmt.Println(prototype.PricingLoop(p))
}
