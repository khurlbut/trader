package api

import (
    "fmt"
    "log"
    "os/exec"
)

func ls() {

    out, err := exec.Command("ls", "-l").Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
}