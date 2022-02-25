package main

import(
	"os"
	"fmt"
	"github.com/khurlbut/trader/api"
	"github.com/khurlbut/trader/campaign"
	"github.com/khurlbut/trader/prototype"
)

func main() {
	api.ls()
	os.Exit(0)
	c := campaign.NewCampaign()
	fmt.Println(prototype.PricingLoop(c))
}
