package cash_percentage_adjuster

import(
  "github.com/magiconair/properties"
)

type CashPercentageTable struct {
	spotPriceLow float64
	spotPriceHigh 	float64
	cashPercentageLow float64
	cashPercentageHigh	float64
}

const propertiesFile = "/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/drive.properties"

func NewCashAdjuster() *CashPercentageTable {
  props := properties.MustLoadFile(propertiesFile, properties.UTF8)

  spotPriceLow := props.GetFloat64("spotPriceLow", 17000.00)
  spotPriceHigh := props.GetFloat64("spotPriceHigh", 25000.00)
  cashPercentageLow := props.GetFloat64("cashPercentageLow", 0.00)
  cashPercentageHigh := props.GetFloat64("cashPercentageHigh", 20.00)

	c := CashPercentageTable{
		spotPriceLow: 	spotPriceLow,
		spotPriceHigh:	spotPriceHigh,
		cashPercentageLow:	cashPercentageLow,	
		cashPercentageHigh:	cashPercentageHigh,
	}
	return &c
}

func (c *CashPercentageTable) CashPercentageTarget(spot float64) float64 {
	spotInRange := findPercentInRange(spot, c.spotPriceLow, c.spotPriceHigh)
	return findPercentInRange(spotInRange, c.cashPercentageLow, c.cashPercentageHigh)
}

func findPercentInRange(value float64, low float64, high float64) float64 {
	return ((value - low) / (high - low))
}
