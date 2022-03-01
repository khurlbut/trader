package main

import(
	"os"
	"fmt"
	"github.com/khurlbut/trader/api"
	"github.com/khurlbut/trader/campaign"
	"github.com/khurlbut/trader/prototype"
	"github.com/magiconair/properties"
)

func main() {
	// props := properties.MustLoadFile("api/binance.properties", properties.UTF8)
	// api_key := props.GetString("api_key_prop", "")
	// secret_key := props.GetString("secret_key_prop", "")
	// fmt.Printf("api_key: %s\nsecret_key: %s\n", api_key, secret_key)
	// // api.Ls()
	// ts := api.Timestamp()
	// api.Order(api_key, secret_key, ts)
	// // fmt.Println(ts)
	// os.Exit(0)
	c := campaign.NewCampaign()
	fmt.Println(prototype.PricingLoop(c))
}
