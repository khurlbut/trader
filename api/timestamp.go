package api

import (
    "fmt"
    "log"
    "os/exec"
)

func Timestamp() {

    out, err := exec.Command("date").Output()
    fmt.Printf("out = %T\n", out)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
    // return fmt.Sprintln("%s000", out')
}