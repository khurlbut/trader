package price_quotes

import (
     "fmt"
     "log"
     "encoding/json"
     // "strings"
     // "strconv"
     "time"
     "github.com/magiconair/properties"
)

type BinanceQuoteService struct {
     baseQuotePair string
     pingEndPoint string
     priceEndPoint string
     currentPrice float64
     pause string
}

type quote struct {
     symbol string
     price float64
}

func NewBinanceQuoteService(propertiesFile string) *BinanceQuoteService {
     props := properties.MustLoadFile(propertiesFile, properties.UTF8)

     baseQuotePair := props.GetString("base_quote_pair", "")
     priceURL := props.GetString("url_price", "")
     queryPrefix := props.GetString("price_query_prefix", "")

     fmt.Printf("baseQuotePair: %s\n", baseQuotePair)
     fmt.Printf("priceURL: %s\n", priceURL)
     fmt.Printf("queryPrefix: %s\n", queryPrefix)

     priceEndPoint := priceURL + queryPrefix + baseQuotePair
     fmt.Printf("priceEndPoint: %s\n", priceEndPoint)

     return &BinanceQuoteService{
          baseQuotePair: props.GetString("base_quote_pair", ""),
          pingEndPoint: props.GetString("url_ping", ""), 
          priceEndPoint: props.GetString("url_price", ""),
          pause: props.GetString("pause", "60s"), 
     }
}

func (qs *BinanceQuoteService) Open() {
     q := "{\"symbol\":\"BTCUSDT\",\"price\":\"37223.53000000\"}"
     err := json.Unmarshal(q, &quote)
     if err != nil {
          log.Fatal(err)
     }
     fmt.Printf("price: %f\n", quote.price)
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