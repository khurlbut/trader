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
     fmt.Println("Open 1")
     f, err := os.Open(qs.datafile)
     if err != nil {
          log.Fatal(err)
     }
     qs.file = f

     fmt.Println("Open 2")
     st, err := time.Parse(qs.dateTimeLayout, qs.startTimeStr) 
     if err != nil {
          log.Fatal(err)
     }
     qs.startTime = st

     fmt.Println("Open 3")
     et, err := time.Parse(qs.dateTimeLayout, qs.endTimeStr) 
     if err != nil {
          log.Fatal(err)
     }
     qs.endTime = et
     
     fmt.Println("Open 4")
     qs.scanner = bufio.NewScanner(file)
     fmt.Println("Open 4.1")
     qs.checkScanner()
     fmt.Println("Open 4.2")
     qs.scanToStartDate()
     fmt.Println("Open 5")
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
     fmt.Println("scatToStartDate 1")
     d := qs.readDate()

     fmt.Println("scatToStartDate 2")
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
     fmt.Println("readDate 1")
     qs.scan()
     fmt.Println("readDate 2")
     t, err := time.Parse(qs.dateTimeLayout, qs.readLineArray()[qs.dateTimeIndex])
     if err != nil {
          log.Fatal(err)
     }
     fmt.Println("readDate 3")
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
     fmt.Println("checkScanner")
     if err := qs.scanner.Err(); err != nil {
          log.Fatal(err)
     }
}
