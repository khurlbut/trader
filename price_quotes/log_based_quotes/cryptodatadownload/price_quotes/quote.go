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
const date_time_index = 1
const date_time_layout = "2006-01-02 15:04:05"
const start_time_str = "2020-11-21 07:21:00"
const end_time_str = "2020-12-21 07:21:00"

var file *os.File = nil
var scanner *bufio.Scanner = nil
var start_time time.Time
var end_time time.Time

func Init() {
     fmt.Println("cryptodatadownload price_quotes Init()")
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
}

func scanToStartDate() {
     // line := readLine()
     
     // t, err := time.Parse(layout, line_arr[date_time_index])
     // if err != nil {
     //      log.Fatal(err)
     // }

     for readDate().Before(start_time) {
         scanner.Scan() 
         checkScanner()
     }
}

func readPrice() float64 {
     p, err := strconv.ParseFloat(readLineArray()[open_index], 64) 
     if err != nil {
          log.Fatal(err)
     }
     return p
}

func readDate() time.Time {
     t, err := time.Parse(layout, readLineArray()[date_time_index])
     if err != nil {
          log.Fatal(err)
     }
     return t
}

func readLineArray() []string {
     l := readLine()
     return strings.Split(line, ",")
}

func readLine() string {
     line := scanner.Text()
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

// var i = 0
func HasNextPrice() bool {
     return readDate().Before(end_time)
     // if i < 100 {
     //      i++
     //      scanner.Scan()
     //      checkScanner()
     //      return true
     // }
     // return false
}

func NextPrice() float64 {
     p := readPrice()
     scanner.Scan()
     return p
     // line := scanner.Text()
     // checkScanner()
     // line_arr := strings.Split(line, ",")
     // p, err := strconv.ParseFloat(line_arr[open_index], 64) 
     // if err != nil {
     //      log.Fatal(err)
     // }
     // // layout := "2006-01-02T15:04:05.000Z"
     // layout := "2006-01-02 15:04:05"
     // t, err := time.Parse(layout, line_arr[date_time_index])
     // if err != nil {
     //      log.Fatal(err)
     // }
     // fmt.Println(t)
     // return p
}

func CurrentPrice() float64 {
     return 0.0
}