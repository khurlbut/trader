package api

import (
    "fmt"
    // "log"
    // "os/exec"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

func Signature(timestamp string, secret_key string, request_body string) string {

    // // out, err := exec.Command("echo", "-n", "timestamp="+timestamp, "|", "openssl", "dgst", "-sha256", "-hmac", secret_key).Output()
    // out, err := exec.Command("`echo -n timestamp="+ timestamp +"| openssl dgst -sha256 -hmac "secret_key"`).Output()

    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println(string(out))

    // secret := "mysecret"
    // data := "data"
    // fmt.Printf("Secret: %s Data: %s\n", secret, data)

    // Create a new HMAC by defining the hash type and the key (as byte array)
    h := hmac.New(sha256.New, []byte(secret_key))

    q := fmt.Sprintf("%s%s", timestamp, request_body)
    // Write Data to it
    h.Write([]byte(q))

    // Get result and encode as hexadecimal string
    sha := hex.EncodeToString(h.Sum(nil))

    // fmt.Println("Result: " + sha)
    fmt.Printf(sha)
    return sha
}