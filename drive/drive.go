package drive

import(
	"github.com/khurlbut/trader/purse"
  "github.com/magiconair/properties"
)

type Drive struct {
	InitialCash float64
	BuyTrigger 	float64
	SellTrigger float64
	Purse *purse.Purse
}

const propertiesFile = "/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/drive.properties"

func NewDrive() *Drive {
  props := properties.MustLoadFile(propertiesFile, properties.UTF8)

  targetCashPercentage := props.GetFloat64("targetCashPercentage", 0.50)
  tradingFeePercentage := props.GetFloat64("tradingFeePercentage", 0.006)
  initialCash := props.GetFloat64("initialCash", 10000.00)
  buyTrigger := props.GetFloat64("buyTrigger", 0.02)
  sellTrigger := props.GetFloat64("sellTrigger", 0.02)

	p := purse.NewPurse(targetCashPercentage, tradingFeePercentage)
	d := Drive{
		InitialCash: 	initialCash,
		BuyTrigger:		buyTrigger,
		SellTrigger:	sellTrigger,
		Purse:				p,
	}
	return &d
}
