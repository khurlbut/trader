package price_quotes

import (
     "fmt"
     "os"
     "log"
     "strings"
     "bufio"
     "strconv"
)

const datafile = "/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/data/cryptodatadownload/Binance_BTCUSDT_minute.csv"
const open_index = 3

var file *os.File = nil
var scanner *bufio.Scanner = nil

func Init() {
     fmt.Println("cryptodatadownload price_quotes Init()")
     f, err := os.Open(datafile)
     if err != nil {
          log.Fatal(err)
     }
     file = f
     scanner = bufio.NewScanner(file)

     checkScanner()
}

func checkScanner() {
     if err := scanner.Err(); err != nil {
          log.Fatal(err)
     }
}

func Close() {
     fmt.Println("cryptodatadownload price_quotes Close()")
     file.Close()
}

var i = 0
func HasNextPrice() bool {
     if i < 100 {
          i++
          scanner.Scan()
          checkScanner()
          return true
     }
     return false
}

func NextPrice() float64 {
     line := scanner.Text()
     checkScanner()
     return strconv.ParseFloat(strings.Split(line, ",")[open_index], 64)
     // fmt.Println(strings.Split(scanner.Text(), ",")[open_index])
     // return 0
}

func CurrentPrice() float64 {
     return 0.0
}