package main

import(
	"os"
	"fmt"
	"github.com/khurlbut/trader/api"
	"github.com/khurlbut/trader/campaign"
	"github.com/khurlbut/trader/prototype"
)

func main() {
	// api.Ls()
	ts := api.Timestamp()
	fmt.Println(ts)
	api.Signature(ts)
	os.Exit(0)
	c := campaign.NewCampaign()
	fmt.Println(prototype.PricingLoop(c))
}
