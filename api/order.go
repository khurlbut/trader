package api

import (
    "fmt"
    "time"
    "strconv"
)

func Order(api_key string, secret_key string, timestamp string, signature string) {
  url := fmt.Sprintf("https://api.binance.us/api/v3/account?timestamp=%s&signature=%s -H X-MBX-APIKEY: %s", timestamp, signature, api_key)
  resp, err := http.Get(qs.priceEndPoint)
  if err != nil {
       log.Fatal(err)
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil {
       log.Fatal(err)
  }  
}