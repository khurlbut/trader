package api

import (
    "fmt"
    "log"
    "os/exec"
)

func Signature(timeStamp string) {

    out, err := exec.Command("echo", "-n", timeStamp).Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
}