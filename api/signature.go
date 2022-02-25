package api

import (
    "fmt"
    "log"
    "os/exec"
)

func Signature(timestamp string) {

    out, err := exec.Command("echo", "-n", "timestamp="+timestamp).Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
}