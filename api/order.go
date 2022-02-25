package api

import (
    "io"
    "fmt"
    "log"
    "net/http"
)

func Order(api_key string, timestamp string, signature string) {
  // url := fmt.Sprintf("https://api.binance.us/api/v3/account?timestamp=%s&signature=%s -H X-MBX-APIKEY: %s", timestamp, signature, api_key)
  url := fmt.Sprintf("https://api.binance.us/api/v3/account?timestamp=%s&signature=%s", timestamp, signature)
  fmt.Printf("\n%s\n",url)
  
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
       log.Fatal(err)
  }
  req.Header.Add("X-MBX-APIKEY", api_key)
  resp, err2 :=client.Do(req)
  if err2 != nil {
       log.Fatal(err)
  }

  defer resp.Body.Close()
  body, err3 := io.ReadAll(resp.Body)
  if err3 != nil {
       log.Fatal(err)
  }  
  fmt.Println(string(body))
}