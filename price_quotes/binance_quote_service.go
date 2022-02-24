package price_quotes

import (
     "fmt"
     "log"
     // "strings"
     // "strconv"
     "time"
     "github.com/magiconair/properties"
)

type BinanceQuoteService struct {
     pause string
}

func NewBinanceQuoteService(propertiesFile string) *BinanceQuoteService {
     props := properties.MustLoadFile(propertiesFile, properties.UTF8)

     pingURL := props.GetString("url_ping", "")
     priceURL := props.GetString("url_price", "")
     queryPrefix := props.GetString("price_query_prefix", "")
     baseQuotePair := props.GetString("base_quote_pair", "")

     fmt.Printf("pingURL: %s\n", pingURL)
     fmt.Printf("priceURL: %s\n", priceURL)
     fmt.Printf("queryPrefix: %s\n", queryPrefix)
     fmt.Printf("baseQuotePair: %s\n", baseQuotePair)

     priceEndPoint := priceURL + queryPrefix + baseQuotePair
     fmt.Printf("priceEndPoint: %s\n", priceEndPoint)

     return &BinanceQuoteService{
          pause: props.GetString("pause", "60s"), 
     }
}

func (qs *BinanceQuoteService) Open() {
}

func (qs *BinanceQuoteService) Close() {
}

func (qs *BinanceQuoteService) HasNextPrice() bool {
     fmt.Println("hasNextPrice")
     return true
}

func (qs *BinanceQuoteService) NextPrice() float64 {
     return -1.0
}

func (qs *BinanceQuoteService) CurrentPrice() float64 {
     return -1.0
}

func (qs *BinanceQuoteService) Pause() {
     s, err := time.ParseDuration(qs.pause)
     if err != nil {
          log.Fatal(err)
     }
     time.Sleep(s)
}