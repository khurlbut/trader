package price_quotes

import (
     "fmt"
     "log"
     "strings"
     "strconv"
     "github.com/magiconair/properties"
)

type StringBasedQuoteService struct {
     spotPriceIndex int
     prices []float64
     currentPrice float64
}

func NewStringBasedQuoteService(propertiesFile string) *StringBasedQuoteService {
     props := properties.MustLoadFile(propertiesFile, properties.UTF8)
     priceStr := props.GetString("prices", "")
     fmt.Printf("priceStr: %s\n", priceStr)

     prices := strings.Split(priceStr, ",")
     priceArr := make([]float64, len(prices))
     for i := range prices {
          price, err := strconv.ParseFloat(strings.TrimSpace(prices[i]), 32)
          if err != nil {
               log.Fatal(err)
          }
          priceArr[i] = price
     }

	return &StringBasedQuoteService{
          spotPriceIndex: 0,
          prices: priceArr,
     }
}

func (qs *StringBasedQuoteService) Open() {
     if HasNextPrice() {
          qs.currentPrice = NextPrice()
     }
}

func (qs *StringBasedQuoteService) Close() {
}

func (qs *StringBasedQuoteService) HasNextPrice() bool {
     return qs.spotPriceIndex < len(qs.prices)
}

func (qs *StringBasedQuoteService) NextPrice() float64 {
     p := qs.prices[qs.spotPriceIndex]
     qs.currentPrice = p
     qs.spotPriceIndex++
     return p
}

func (qs *StringBasedQuoteService) CurrentPrice() float64 {
     return qs.currentPrice
}