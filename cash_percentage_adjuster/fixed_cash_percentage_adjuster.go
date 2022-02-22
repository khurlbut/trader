package cash_percentage_adjuster

import(
	"log"
  // "github.com/magiconair/properties"
)

type FixedCashPercentageTable struct {}

// const propertiesFile = "/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/campaign.properties"

func NewFixedCashAdjuster() *FixedCashPercentageTable {
	log.Fatal("Not Implemented")
  // props := properties.MustLoadFile(propertiesFile, properties.UTF8)
	// return nil
}

func (c *FixedCashPercentageTable) CashPercentageTarget(spot float64) float64 {
	log.Fatal("Not Implemented")
	return -1.0
}
