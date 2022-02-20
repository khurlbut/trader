package main

import(
	"fmt"
	"github.com/khurlbut/trader/prototype"
	"github.com/khurlbut/trader/purse"
)

func main() {
	var p *purse.Purse
	p = purse.NewPurse(123.456, 789.123)
	fmt.println("p: " + p)
	fmt.Println(prototype.PricingLoop())
}
