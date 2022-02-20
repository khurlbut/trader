package main

import(
	"fmt"
	"github.com/khurlbut/trader/prototype"
	"github.com/khurlbut/trader/purse"
)

func main() {
p := purse.NewPurse(0.5, 0.006)
	p.FundPurse(10000, 1000)
	fmt.Printf("p.Properties: %s\n", p.Properties())
	fmt.Printf("p.Holdings: %s\n", p.Holdings(1000))
	fmt.Println(prototype.PricingLoop())
}
