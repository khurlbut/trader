package api

import (
    "io"
    "fmt"
    "log"
    "net/http"
    "bytes"
)

func Order(api_key string, secret_key string, timestamp string) {
  // url := fmt.Sprintf("https://api.binance.us/api/v3/account?timestamp=%s&signature=%s -H X-MBX-APIKEY: %s", timestamp, signature, api_key)
  // url := fmt.Sprintf("https://api.binance.us/api/v3/account?timestamp=%s&signature=%s", timestamp, signature)
  url := "https://api.binance.us/api/v3/account"
  fmt.Printf("\n%s\n",url)
  
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
       log.Fatal(err)
  }
  buf := new(bytes.Buffer)
  buf.ReadFrom(req.GetBody()) 
  signature := Signature(timestamp, secret_key, buf.String())
  fmt.Println("signature: ", signature)
  q := req.URL.Query()
  q.Add("signature", signature)
  q.Add("timestamp", timestamp)
  req.URL.RawQuery = q.Encode()

  req.Header.Add("X-MBX-APIKEY", api_key)

  fmt.Println(req.URL.String())
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