package main

import(
	"os"
	"fmt"
	"github.com/khurlbut/trader/api"
	"github.com/khurlbut/trader/campaign"
	"github.com/khurlbut/trader/prototype"
	// "github.com/magiconair/properties"
)

func main() {
	// props := properties.MustLoadFile("api/binance.properties", properties.UTF8)
	// api.Ls()
	ts := api.Timestamp()
	// fmt.Println(ts)
	api.Signature(ts, "YT4DtosTWcptwdlScXUISNA9FQ5FxsXGtnSoROeTxyuJBQEGfi8X9lziav2bW9tf")
	os.Exit(0)
	c := campaign.NewCampaign()
	fmt.Println(prototype.PricingLoop(c))
}
