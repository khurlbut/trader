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

     ping_url := props.GetString("url_ping", "")
     price_url := props.GetString("url_price", "")

     fmt.Printf("ping_url: %s\n", ping_url)
     fmt.Printf("price_url: %s\n", price_url)

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