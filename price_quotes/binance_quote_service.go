package price_quotes

import (
     "fmt"
     "log"
     "encoding/json"
     "strings"
     "strconv"
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
     Symbol string
     Price string
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
     var r = []byte(`{"symbol":"BTCUSDT","price":"37223.53000000"}`)
     // qs.currentPrice = qs.readPrice(r)
     qs.currentPrice = qs.readPrice()

     // fmt.Printf("quote: %+v\n", quote)
     // fmt.Printf("symbol: %s\n", quote.Symbol)
     fmt.Printf("price: %f\n", qs.currentPrice)
}

// func (qs *BinanaceQuoteService) readPrice(bytes []bytes) float64 {
func (qs *BinanaceQuoteService) readPrice() float64 {
     // q := unmarshal(bytes)
     // if q.Symbol != qs.baseQuotePair {
     //      log.Fatal("Pair mismatch!")
     // }
     // return parseFloat(q.Price)
     return 0.0
}

func unmarshal(bytes []byte) *quote {
     var q = &quote{}
     err := json.Unmarshal(bytes, q)
     if err != nil {
          log.Fatal(err)
     }
     return q
}

func parseFloat(s string) float64 {
     p, err := strconv.ParseFloat(strings.TrimSpace(s), 32)
     if err != nil {
          log.Fatal(err)
     }
     return p
}

func (qs *BinanceQuoteService) Close() {
}

func (qs *BinanceQuoteService) HasNextPrice() bool {
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