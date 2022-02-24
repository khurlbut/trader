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

     return &BinanceQuoteService{
          baseQuotePair: props.GetString("base_quote_pair", ""),
          pingEndPoint: props.GetString("url_ping", ""), 
          priceEndPoint: buildPriceURL(props),
          pause: props.GetString("pause", "60s"), 
     }
}

func (qs *BinanceQuoteService) Open() {
     qs.currentPrice = qs.readPrice(qs.httpGetPriceQuote())
}

func (qs *BinanceQuoteService) Close() {
}

func (qs *BinanceQuoteService) HasNextPrice() bool {
     return false
}

func (qs *BinanceQuoteService) NextPrice() float64 {
     return qs.currentPrice
}

func (qs *BinanceQuoteService) CurrentPrice() float64 {
     return qs.currentPrice
}

func (qs *BinanceQuoteService) Pause() {
     s, err := time.ParseDuration(qs.pause)
     if err != nil {
          log.Fatal(err)
     }
     time.Sleep(s)
}

func (qs *BinanceQuoteService) httpGetPriceQuote() []byte {
     // resp, err := http.Get("http://example.com/")
     resp, err := http.Get(qs.priceEndPoint)
     if err != nil {
          // handle error
     }
     defer resp.Body.Close()
     body, err := io.ReadAll(resp.Body)
     // return []byte(`{"symbol":"BTCUSDT","price":"37223.53000000"}`)
     return io.ReadAll(resp.Body)
}

func (qs *BinanceQuoteService) readPrice(bytes []byte) float64 {
     q := unmarshal(bytes)
     if q.Symbol != qs.baseQuotePair {
          log.Fatal("Pair mismatch!")
     }
     return parseFloat(q.Price)
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

func buildPriceURL(props *properties.Properties) string {
     baseQuotePair := props.GetString("base_quote_pair", "")
     priceURL := props.GetString("url_price", "")
     queryPrefix := props.GetString("price_query_prefix", "")
     priceEndPoint := priceURL + queryPrefix + baseQuotePair
     fmt.Printf("priceEndPoint: %s\n", priceEndPoint)

     return priceEndPoint
}
