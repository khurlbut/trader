package main

import(
	"fmt"
	"github.com/khurlbut/trader/prototype"
	"github.com/khurlbut/trader/purse"
)

func main() {
	var purse *Purse
	purse = NewPurse(123.456, 789.123)
	fmt.Println(prototype.PricingLoop())
}
