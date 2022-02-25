package api

import (
    "fmt"
    "log"
    "os/exec"
)

func Signature(timestamp string, secret_key string) {

    out, err := exec.Command("echo", "-n", "timestamp="+timestamp, "|", "openssl", "dgst", "-sha256", "-hmac", secret_key).Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
}