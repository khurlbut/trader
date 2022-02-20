package main

import(
	"fmt"
	"github.com/khurlbut/trader/prototype"
	"github.com/khurlbut/trader/purse"
)

func main() {
	p := purse.NewPurse(123.456, 789.123)
	p.FundPurse(10000, 1000)
	fmt.Printf("p.Properties: %s\n", p.Properties())
	fmt.Printf("p.Holdings: %s\n", p.Holdings())
	// fmt.Println(prototype.PricingLoop())
}
