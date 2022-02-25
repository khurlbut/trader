package main

import(
	"os"
	"fmt"
	"github.com/khurlbut/trader/campaign"
	"github.com/khurlbut/trader/prototype"
)

func main() {
	os.Exit(0)
	c := campaign.NewCampaign()
	fmt.Println(prototype.PricingLoop(c))
}
