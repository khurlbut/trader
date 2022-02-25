package campaign

import(
	"log"
	"github.com/khurlbut/trader/price_quotes"
	"github.com/khurlbut/trader/purse"
  "github.com/magiconair/properties"
)

type Campaign struct {
	QuoteService price_quotes.QuoteService
	InitialCash float64
	BuyTrigger 	float64
	SellTrigger float64
	Purse *purse.Purse
}

const propertiesFile = "campaign.properties"

func NewCampaign() *Campaign {
  props := properties.MustLoadFile(propertiesFile, properties.UTF8)

	baseQuotePair := props.GetString("base_quote_pair", "")
	quoteService := props.GetString("quoteService", "")
	quoteServicePropsFile := props.GetString("quoteServicePropsFile", "")

	var qs price_quotes.QuoteService = nil
	if quoteService == "BinanceQuoteService" {
		qs = price_quotes.NewBinanceQuoteService(baseQuotePair, quoteServicePropsFile)
	} else if quoteService == "StringBasedQuoteService" {
		qs = price_quotes.NewStringBasedQuoteService(quoteServicePropsFile)
	} else if quoteService == "CommaSeparatedValueQuoteService" {
		qs = price_quotes.NewCommaSeparatedValueQuoteService(quoteServicePropsFile)
	} else {
		log.Fatal("Invalid Quote Service: " + quoteService)
	}

  targetCashPercentage := props.GetFloat64("targetCashPercentage", 0.00)
  tradingFeePercentage := props.GetFloat64("tradingFeePercentage", 0.000)
  initialCash := props.GetFloat64("initialCash", 0.00)
  buyTrigger := props.GetFloat64("buyTrigger", 0.00)
  sellTrigger := props.GetFloat64("sellTrigger", 0.00)

	p := purse.NewPurse(targetCashPercentage, tradingFeePercentage)
	c := Campaign{
		QuoteService: qs,
		InitialCash: 	initialCash,
		BuyTrigger:		buyTrigger,
		SellTrigger:	sellTrigger,
		Purse:				p,
	}
	return &c
}
