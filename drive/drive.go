package drive

import(
  "fmt"
	"github.com/khurlbut/trader/purse"
  "github.com/magiconair/properties"
)

// const targetCashPercentage = 0.5
const tradingFeePercentage = 0.006

type Drive struct {
	InitialCash float64
	BuyTrigger 	float64
	SellTrigger float64
	Purse *purse.Purse
}

// func NewDrive(initialCash float64, buyTrigger float64, SellTrigger float64) *Drive {
func NewDrive() *Drive {

  props := properties.MustLoadFile("/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/drive.properties", properties.UTF8)
  targetCashPercentage := props.GetFloat64("targetCashPercentage", 0.50)
  tradingFeePercentage := props.GetFloat64("tradingFeePercentage", 0.006)
  initialCash := props.GetFloat64("initialCash", 10000.00)
  buyTrigger := props.GetFloat64("buyTrigger", 0.02)
  sellTrigger := props.GetFloat64("sellTrigger", 0.02)

	p := purse.NewPurse(targetCashPercentage, tradingFeePercentage)
	d := Drive{
		InitialCash: 	10000.00,
		BuyTrigger:		0.02,
		SellTrigger:	0.02,
		Purse:				p,
	}
	return &d
}
