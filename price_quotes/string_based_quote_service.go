package price_quotes

import (
     "strings"
     "strconv"
     "github.com/magiconair/properties"
)

type StringBasedQuoteService struct {
     spotPriceIndex int
     prices []float64
}

func NewStringBasedQuoteService(propertiesFile string) *StringBasedQuoteService {
     props := properties.MustLoadFile(propertiesFile, properties.UTF8)
     priceStr := props.GetString("prices", "")

     prices := strings.Split(priceStr, ",")
     arr := make([]float64, len(prices))
     for i := range arr {
          arr[i], _ = strconv.ParseFloat(prices[i], 32)
     }

	return &StringBasedQuoteService{
          spotPriceIndex: 0,
          prices: arr,
     }
}

func (qs *StringBasedQuoteService) Open() {
}

func (qs *StringBasedQuoteService) Close() {
}

func (qs *StringBasedQuoteService) HasNextPrice() bool {
     return qs.spotPriceIndex < len(qs.prices)
}

func (qs *StringBasedQuoteService) NextPrice() float64 {
     p := qs.prices[qs.spotPriceIndex]
     qs.spotPriceIndex++
     return p
}

func (qs *StringBasedQuoteService) CurrentPrice() float64 {
     return qs.prices[qs.spotPriceIndex]
}