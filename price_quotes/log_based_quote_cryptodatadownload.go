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

func (qs *CommaSeparatedValueQuoteService) Open() {
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
     
     qs.scanner = bufio.NewScanner(qs.file)
     qs.checkScanner()
     qs.scanToStartDate()
}

func (qs *CommaSeparatedValueQuoteService) HasNextPrice() bool {
     d := qs.readDate()
     if d.Before(qs.endTime) {
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

     for d.Before(qs.startTime) {
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
