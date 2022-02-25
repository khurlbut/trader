package api

import (
    "fmt"
    "log"
    "os/exec"
)

func Timestamp() {

    out, err := exec.Command("date +v%s000").Output()
    fmt.Printf("out = %T\n", out)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
    // return fmt.Sprintln("%s000", out')
}