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


type CryptoDataDownloadQuoteService struct {}

func NewCryptoDataDownloadQuoteService() *CryptoDataDownloadQuoteService {
	return &CryptoDataDownloadQuoteService{}
}

const datafile = "/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/data/cryptodatadownload/Binance_BTCUSDT_minute.csv"
const open_index = 3
const date_time_index = 1
const date_time_layout = "2006-01-02 15:04:05"
const start_time_str = "2020-11-21 07:21:00"
const end_time_str = "2020-12-21 07:21:00"

var file *os.File = nil
var scanner *bufio.Scanner = nil
var start_time time.Time
var end_time time.Time

func (CryptoDataDownloadQuoteService) Open() {
     f, err := os.Open(datafile)
     if err != nil {
          log.Fatal(err)
     }
     file = f

     start_time, err = time.Parse(date_time_layout, start_time_str) 
     if err != nil {
          log.Fatal(err)
     }

     end_time, err = time.Parse(date_time_layout, end_time_str) 
     if err != nil {
          log.Fatal(err)
     }
     
     scanner = bufio.NewScanner(file)
     checkScanner()
     scanToStartDate()
}

func (CryptoDataDownloadQuoteService) HasNextPrice() bool {
     d := readDate()
     if d.Before(end_time) {
          return true
     }
     fmt.Println("End Time: " + d.String())
     return false
}

func (CryptoDataDownloadQuoteService) NextPrice() float64 {
     p := readPrice()
     scan()
     return p
}

func (CryptoDataDownloadQuoteService) CurrentPrice() float64 {
     return readPrice()
}

func (CryptoDataDownloadQuoteService) Close() {
     file.Close()
}

func scanToStartDate() {
     d := readDate()

     for d.Before(start_time) {
          d = readDate()
     }
     fmt.Println("Start Time: " + d.String())
}

func readPrice() float64 {
     p, err := strconv.ParseFloat(readLineArray()[open_index], 64) 
     if err != nil {
          log.Fatal(err)
     }
     return p
}

func readDate() time.Time {
     scan()
     t, err := time.Parse(date_time_layout, readLineArray()[date_time_index])
     if err != nil {
          log.Fatal(err)
     }
     return t
}

func readLineArray() []string {
     return strings.Split(readLine(), ",")
}

func readLine() string {
     l := scanner.Text()
     checkScanner()
     return l
}

func scan() {
     scanner.Scan()
     checkScanner()
}

func checkScanner() {
     if err := scanner.Err(); err != nil {
          log.Fatal(err)
     }
}