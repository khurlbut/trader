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
  
  client := &http.Client{
    CheckRedirect: redirectPolicyFunc,
  }
  req, err := httpNewRequest("GET", url, nil)
  req.Header.Add("X-MBX-APIKEY", api_key)

  resp, err :=client.Do(req)
  if err != nil {
       log.Fatal(err)
  }

  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil {
       log.Fatal(err)
  }  
  fmt.Println(string(body))
}