package price_quotes

import (
     "fmt"
     "os"
     "log"
)

// var spotPriceIndex = 0
// var prices = []float64{10.00, 10.200, 10.404, 10.19592, 9.98784, 10.00}

var file File = nil

func Init() {
     fmt.Println("cryptodatadownload price_quotes Init()")
     file, err := os.Open("~/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/data/cryptodatadownload/Binance_BTCUSDT_minute.csv")
     if err != nil {
          log.Fatal(err)
     }
}

func Close() {
     fmt.Println("cryptodatadownload price_quotes Close()")
     file.Close()
}

func HasNextPrice() bool {
     false
     // return spotPriceIndex < len(prices)
}

func NextPrice() float64 {
     return nil
     // p := prices[spotPriceIndex]
     // spotPriceIndex++
     // return p
}

func CurrentPrice() float64 {
     return 0.0
     // return prices[spotPriceIndex]
}