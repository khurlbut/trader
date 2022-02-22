package price_quotes

import (
     "fmt"
     "os"
     "log"
     "strings"
     "bufio"
     "strconv"
     "time"
     "github.com/magiconair/properties"
)


type CommaSeparatedValueQuoteService struct {
     file *os.File 
     scanner *bufio.Scanner
     datafile string
     spotPriceIndex int
     dateTimeIndex int
     dateTimeLayout string
     startTimeStr string
     endTimeStr string
     startTime time.Time
     endTime time.Time
}

func NewCommaSeparatedValueQuoteService(propertiesFile string)*CommaSeparatedValueQuoteService {
     propertiesPath := "price_quotes/"
     props := properties.MustLoadFile(propertiesPath + propertiesFile, properties.UTF8)

	return &CommaSeparatedValueQuoteService{
          datafile: props.GetString("datafile", ""),
          spotPriceIndex: props.GetInt("spot_price_index", -1),
          dateTimeIndex: props.GetInt("date_time_index", -1),
          dateTimeLayout: props.GetString("date_time_layout", ""),
          startTimeStr: props.GetString("start_time", ""),
          endTimeStr: props.GetString("end_time", ""),
     }
}

// const spot_price_index = 3
// const date_time_index = 1
// const date_time_layout = "2006-01-02 15:04:05"

// const datafile = "/Users/Ke015t7/.gvm/pkgsets/go1.17.7/global/src/github.com/khurlbut/trader/data/cryptodatadownload/Binance_BTCUSDT_minute.csv"
// const start_time_str = "2020-11-21 07:21:00"
// const end_time_str = "2020-12-21 07:21:00"

var file *os.File = nil
var scanner *bufio.Scanner = nil
var start_time time.Time
var end_time time.Time

func (qs *CommaSeparatedValueQuoteService) Open() {
     fmt.Println("Open")
     f, err := os.Open(qs.datafile)
     if err != nil {
          log.Fatal(err)
     }
     qs.file = f

     st, err := time.Parse(qs.dateTimeLayout, qs.startTimeStr) 
     if err != nil {
          log.Fatal(err)
     }
     qs.startTime = st

     et, err := time.Parse(qs.dateTimeLayout, qs.endTimeStr) 
     if err != nil {
          log.Fatal(err)
     }
     qs.endTime = et
     
     qs.scanner = bufio.NewScanner(file)
     qs.checkScanner()
     qs.scanToStartDate()
}

func (qs *CommaSeparatedValueQuoteService) HasNextPrice() bool {
     d := qs.readDate()
     if d.Before(end_time) {
          return true
     }
     fmt.Println("End Time: " + d.String())
     return false
}

func (qs *CommaSeparatedValueQuoteService) NextPrice() float64 {
     p := qs.readPrice()
     qs.scan()
     return p
}

func (qs *CommaSeparatedValueQuoteService) CurrentPrice() float64 {
     return qs.readPrice()
}

func (qs *CommaSeparatedValueQuoteService) Close() {
     qs.file.Close()
}

func (qs *CommaSeparatedValueQuoteService) scanToStartDate() {
     d := qs.readDate()

     for d.Before(start_time) {
          d = qs.readDate()
     }
     fmt.Println("Start Time: " + d.String())
}

func (qs *CommaSeparatedValueQuoteService) readPrice() float64 {
     p, err := strconv.ParseFloat(qs.readLineArray()[qs.spotPriceIndex], 64) 
     if err != nil {
          log.Fatal(err)
     }
     return p
}

func (qs *CommaSeparatedValueQuoteService) readDate() time.Time {
     qs.scan()
     t, err := time.Parse(qs.dateTimeLayout, qs.readLineArray()[qs.dateTimeIndex])
     if err != nil {
          log.Fatal(err)
     }
     return t
}

func (qs *CommaSeparatedValueQuoteService) readLineArray() []string {
     return strings.Split(qs.readLine(), ",")
}

func (qs *CommaSeparatedValueQuoteService) readLine() string {
     l := qs.scanner.Text()
     qs.checkScanner()
     return l
}

func (qs *CommaSeparatedValueQuoteService) scan() {
     qs.scanner.Scan()
     qs.checkScanner()
}

func (qs *CommaSeparatedValueQuoteService) checkScanner() {
     if err := qs.scanner.Err(); err != nil {
          log.Fatal(err)
     }
}
