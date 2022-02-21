package main

import(
	"fmt"
	"github.com/khurlbut/trader/drive"
	"github.com/khurlbut/trader/prototype"
)

// const targetCashPercentage = 0.5
// const tradingFeePercentage = 0.006

// type Drive struct {
// 	InitialCash float64
// 	BuyTrigger 	float64
// 	SellTrigger float64
// 	Purse *purse.Purse
// }

// func NewDrive(initialCash float64, buyTrigger float64, SellTrigger float64) *Drive {
// 	p := purse.NewPurse(targetCashPercentage, tradingFeePercentage)
// 	d := Drive{
// 		InitialCash: 	10000.00,
// 		BuyTrigger:		0.02,
// 		SellTrigger:	0.02,
// 		Purse:				p,
// 	}
// 	return &d
// }

func main() {
	// p := purse.NewPurse(targetCashPercentage, tradingFeePercentage)
	d := drive.NewDrive()
	fmt.Println(prototype.PricingLoop(d))
}
