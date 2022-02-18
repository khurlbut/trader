package price_quotes

import (
     "fmt"
     "os"
     "log"
     "strings"
     "bufio"
     "strconv"
     "time"
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
     line_arr := strings.Split(line, ",")
     p, err := strconv.ParseFloat(line_arr[open_index], 64) 
     if err != nil {
          log.Fatal(err)
     }
     layout := "2006-01-02T15:04:05.000Z"
     t, err := time.Parse(layout, line_arr[1])
     if err != nil {
          log.Fatal(err)
     }
     fmt.Println(t)
     return p
}

func CurrentPrice() float64 {
     return 0.0
}