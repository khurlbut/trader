package price_quotes

import (
     "fmt"
     "os"
     "log"
)

const datafile = "/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/data/cryptodatadownload/Binance_BTCUSDT_minute.csv"

var file *os.File = nil

func Init() {
     fmt.Println("cryptodatadownload price_quotes Init()")
     f, err := os.Open(datafile)
     if err != nil {
          log.Fatal(err)
     }
     file = f
}

func Close() {
     fmt.Println("cryptodatadownload price_quotes Close()")
     file.Close()
}

var i = 0
func HasNextPrice() bool {
     if i < 10 {
          i++
          return true
     }
     return false
}

func NextPrice() float64 {
     return 0
}

func CurrentPrice() float64 {
     return 0.0
}