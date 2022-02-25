package api

import (
    "io"
    "fmt"
    "log"
    "net/http"
)

func Order(api_key string, timestamp string, signature string) {
  url := fmt.Sprintf("https://api.binance.us/api/v3/account?timestamp=%s&signature=%s -H X-MBX-APIKEY: %s", timestamp, signature, api_key)
  fmt.Printf("\n%s\n",url)
  resp, err := http.Get(url)
  if err != nil {
       log.Fatal(err)
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil {
       log.Fatal(err)
  }  
  fmt.Println(body)
}