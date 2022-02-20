package main

import(
	"fmt"
	"github.com/khurlbut/trader/prototype"
	"github.com/khurlbut/trader/purse"
)

func main() {
	p := purse.NewPurse(123.456, 789.123)
	fmt.Printf("p.NewString: %" + p.NewString())
	fmt.Println(prototype.PricingLoop())
}
