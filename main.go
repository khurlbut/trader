package main

import(
	"fmt"
	"github.com/khurlbut/trader/campaign"
	"github.com/khurlbut/trader/prototype"
)

func main() {
	c := drive.NewCampaign()
	fmt.Println(prototype.PricingLoop(c))
}
