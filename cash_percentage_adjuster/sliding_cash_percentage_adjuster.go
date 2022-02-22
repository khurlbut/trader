package cash_percentage_adjuster

import(
  "github.com/magiconair/properties"
)

type SlidingCashPercentageTable struct {
	spotPriceLow float64
	spotPriceHigh 	float64
	cashPercentageLow float64
	cashPercentageHigh	float64
}

const propertiesFile = "/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/campaign.properties"

func NewSlidingCashAdjuster() *SlidingCashPercentageTable {
  props := properties.MustLoadFile(propertiesFile, properties.UTF8)

  spotPriceLow := props.GetFloat64("spotPriceLow", 0.00)
  spotPriceHigh := props.GetFloat64("spotPriceHigh", 0.00)
  cashPercentageLow := props.GetFloat64("cashPercentageLow", 0.00)
  cashPercentageHigh := props.GetFloat64("cashPercentageHigh", 0.00)

	c := SlidingCashPercentageTable{
		spotPriceLow: 	spotPriceLow,
		spotPriceHigh:	spotPriceHigh,
		cashPercentageLow:	cashPercentageLow,	
		cashPercentageHigh:	cashPercentageHigh,
	}
	return &c
}

func (c *SlidingCashPercentageTable) CashPercentageTarget(spot float64) float64 {
	if spot <= c.spotPriceLow {
		return c.cashPercentageLow
	}
	if (spot >= c.spotPriceHigh) {
		return c.cashPercentageHigh
	}

	spotInRange := findPercentInRange(spot, c.spotPriceLow, c.spotPriceHigh)
	return spotInRange * (c.cashPercentageHigh - c.cashPercentageLow) + c.cashPercentageLow
}

func findPercentInRange(value float64, low float64, high float64) float64 {
	return ((value - low) / (high - low))
}
